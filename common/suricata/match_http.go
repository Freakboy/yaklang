package suricata

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"github.com/yaklang/yaklang/common/log"
	"github.com/yaklang/yaklang/common/utils/lowhttp"
	"io"
)

type HTTPConfig struct {
	// deprecated and not implemented
	Uricontent string

	// not set 0
	// equal 1
	// bigger than 2
	// smaller than 3
	// between 4
	UrilenOp   int
	UrilenNum1 int
	UrilenNum2 int
}

func httpMatcher(c *matchContext) error {
	if c.Rule.ContentRuleConfig == nil {
		return nil
	}

	httpraw := c.PK.ApplicationLayer().LayerContents()
	if len(httpraw) == 0 {
		return nil
	}

	if lowhttp.IsResp(httpraw) {
		return httpResMatcher(c)
	}
	return httpReqMatcher(c)
}

func httpReqMatcher(c *matchContext) error {
	payload := c.PK.ApplicationLayer().LayerContents()
	request, err := lowhttp.ParseBytesToHttpRequest(payload)
	if !c.Must(err == nil) {
		//"parse httpraw as http request failed"
		log.Debugf("parse httpraw as http request failed: %v", err)
		return nil
	}

	if cf := c.Rule.ContentRuleConfig.HTTPConfig; cf != nil {
		if cf.Uricontent != "" {
			log.Errorf("uricontent has been deprecated and not implemented yet")
		}
		switch cf.UrilenOp {
		case 1:
			if !c.Must(c.Rule.ContentRuleConfig.HTTPConfig.UrilenNum1 == len(request.URL.Path)) {
				return nil
			}
		case 2:
			if !c.Must(c.Rule.ContentRuleConfig.HTTPConfig.UrilenNum1 > len(request.URL.Path)) {
				return nil
			}
		case 3:
			if !c.Must(c.Rule.ContentRuleConfig.HTTPConfig.UrilenNum1 < len(request.URL.Path)) {
				return nil
			}
		case 4:
			if !c.Must(c.Rule.ContentRuleConfig.HTTPConfig.UrilenNum1 < len(request.URL.Path) && c.Rule.ContentRuleConfig.HTTPConfig.UrilenNum2 > len(request.URL.Path)) {
				return nil
			}
		default:
			// not set
		}
	}

	for _, r := range c.Rule.ContentRuleConfig.ContentRules {
		switch r.Modifier {
		case HTTPUri:
			c.Attach(newPayloadMatcher(r, []byte(request.URL.Path)))
		case HTTPUriRaw:
			rd := bufio.NewReader(bytes.NewReader(payload))
			if _, err := rd.ReadBytes(' '); !c.Must(err == nil) {
				return nil
			}
			uriraw, err := rd.ReadBytes(' ')
			if !c.Must(err == nil) {
				return nil
			}
			if !c.Must(len(uriraw) != 0) {
				return nil
			}
			c.Attach(newPayloadMatcher(r, uriraw[:len(uriraw)-1]))
		case HTTPMethod:
			c.Attach(newPayloadMatcher(r, []byte(request.Method)))
		case HTTPRequestLine:
			idx := bytes.Index(payload, []byte("\r\n"))
			if !c.Must(idx != -1) {
				return nil
			}
			c.Attach(newPayloadMatcher(r, payload[:idx+2]))
		case HTTPRequestBody:
			all, err := io.ReadAll(request.Body)
			if !c.Must(err == nil) {
				return nil
			}
			c.Attach(newPayloadMatcher(r, all))
		case HTTPUserAgent:
			c.Attach(newPayloadMatcher(r, []byte(request.UserAgent())))
		case HTTPHost:
			c.Attach(newPayloadMatcher(r, []byte(request.Host)))
		case HTTPHostRaw:
			st := bytes.Index(payload, []byte("\r\nHost: "))
			if !c.Must(st != -1) {
				return nil
			}
			ed := bytes.Index(payload[st+8:], []byte("\r\n"))
			if !c.Must(ed != -1) {
				return nil
			}
			c.Attach(newPayloadMatcher(r, payload[st+8:st+8+ed]))
		case HTTPAccept:
			c.Attach(newPayloadMatcher(r, []byte(request.Header.Get("Accept"))))
		case HTTPAcceptLang:
			c.Attach(newPayloadMatcher(r, []byte(request.Header.Get("Accept-Language"))))
		case HTTPAcceptEnc:
			c.Attach(newPayloadMatcher(r, []byte(request.Header.Get("Accept-Encoding"))))
		case HTTPReferer:
			c.Attach(newPayloadMatcher(r, []byte(request.Header.Get("Referer"))))
		case HTTPHeader:
			var bb bytes.Buffer
			for k, v := range request.Header {
				bb.WriteString(k)
				bb.WriteString(": ")
				bb.WriteString(v[0])
				for i := 1; i < len(v); i++ {
					bb.WriteString(", ")
					bb.WriteString(v[i])
				}
				bb.WriteString("\r\n")
			}
			c.Attach(newPayloadMatcher(r, bb.Bytes()))
		case HTTPHeaderRaw:
			ed := bytes.Index(payload, []byte("\r\n\r\n"))
			st := bytes.Index(payload, []byte("\r\n"))
			if !c.Must(ed-st > 2) {
				return nil
			}
			c.Attach(newPayloadMatcher(r, payload[st+2:ed+2]))
		case HTTPCookie:
			c.Attach(newPayloadMatcher(r, []byte(request.Header.Get("Cookie"))))
		case HTTPConnection:
			c.Attach(newPayloadMatcher(r, []byte(request.Header.Get("Connection"))))
		case FileData:
			c.Attach(newFileDataMatcher(r, request))
		case HTTPContentType:
			c.Attach(newPayloadMatcher(r, []byte(request.Header.Get("Content-Type"))))
		case HTTPContentLen:
			c.Attach(newPayloadMatcher(r, []byte(request.Header.Get("Content-Length"))))
		case HTTPStart:
			idx := bytes.Index(payload, []byte("\r\n\r\n"))
			if !c.Must(idx != -1) {
				return nil
			}
			c.Attach(newPayloadMatcher(r, lowhttp.FixHTTPRequestOut(payload[:idx+4])))
		case HTTPProtocol:
			c.Attach(newPayloadMatcher(r, []byte(request.Proto)))
		case HTTPHeaderNames:
			var bb bytes.Buffer
			for k := range request.Header {
				bb.WriteString("\r\n")
				bb.WriteString(k)
			}
			bb.WriteString("\r\n\r\n")
			c.Attach(newPayloadMatcher(r, bb.Bytes()))
		}
		err := c.Next()
		if !c.Must(err == nil) {
			return err
		}
		if c.rejected {
			return nil
		}
	}
	return nil
}

func httpResMatcher(c *matchContext) error {
	payload := c.PK.ApplicationLayer().LayerContents()
	response, err := lowhttp.ParseBytesToHttpRequest(payload)
	if err != nil {
		return errors.Wrap(err, "parse httpraw as http response failed")
	}
	_ = response
	return nil

}
