package yakgrpc

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/yaklang/yaklang/common/consts"
	"github.com/yaklang/yaklang/common/log"
	"github.com/yaklang/yaklang/common/utils"
	"github.com/yaklang/yaklang/common/utils/lowhttp"
	"github.com/yaklang/yaklang/common/yak"
	"github.com/yaklang/yaklang/common/yakgrpc/yakit"
	"github.com/yaklang/yaklang/common/yakgrpc/ypb"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestH2Hijack(t *testing.T) {
	h2Host, h2Port := utils.DebugMockHTTP2(context.Background(), func(req []byte) []byte {
		return req
	})
	h2Addr := utils.HostPort(h2Host, h2Port)
	_, err := yak.NewScriptEngine(10).ExecuteEx(`
rsp,req = poc.HTTP(getParam("packet"), poc.http2(true), poc.https(true))~
dump(rsp)
dump(req)
`, map[string]any{
		"packet": `GET / HTTP/2.0
User-Agent: 111
Host: ` + h2Addr,
	})
	if err != nil {
		panic(err)
	}
}

func TestGRPCMUSTPASS_MITM(t *testing.T) {
	client, err := NewLocalClient()
	if err != nil {
		panic(err)
	}

	var (
		started           bool
		passthroughTested bool
		echoTested        bool
		gzipAutoDecode    bool
		chunkDecode       bool
		h2Test            bool
	)

	var mockHost, mockPort = utils.DebugMockHTTPEx(func(req []byte) []byte {
		passthroughTested = true
		rsp, _, _ := lowhttp.FixHTTPResponse([]byte(`HTTP/1.1 200 OK\n
Content-Type: text/html
Content-Length: 3

111`))
		_, body := lowhttp.SplitHTTPHeadersAndBodyFromPacket(req)
		rsp = lowhttp.ReplaceHTTPPacketBodyFast(rsp, body)
		if lowhttp.GetHTTPPacketHeader(req, "Content-Encoding") == "gzip" {
			rsp = lowhttp.ReplaceHTTPPacketHeader(rsp, "Content-Encoding", "gzip")
		}
		if lowhttp.GetHTTPPacketHeader(req, "Transfer-Encoding") == "chunked" {
			rsp = lowhttp.ReplaceHTTPPacketHeader(rsp, "Transfer-Encoding", "chunked")
		}
		return rsp
	})

	log.Infof("start to mock server: %v", utils.HostPort(mockHost, mockPort))
	var rPort = utils.GetRandomAvailableTCPPort()
	var proxy = "http://127.0.0.1:" + fmt.Sprint(rPort)
	_ = proxy

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	/* H2 */
	h2Host, h2Port := utils.DebugMockHTTP2(ctx, func(req []byte) []byte {
		return req
	})
	h2Addr := utils.HostPort(h2Host, h2Port)
	_, err = yak.NewScriptEngine(10).ExecuteEx(`
rsp,req = poc.HTTP(getParam("packet"), poc.http2(true), poc.https(true))~
dump(rsp)
dump(req)
`, map[string]any{
		"packet": `GET / HTTP/2.0
User-Agent: 111
Host: ` + h2Addr,
	})
	if err != nil {
		panic(err)
	}

	stream, err := client.MITM(ctx)
	if err != nil {
		panic(err)
	}
	stream.Send(&ypb.MITMRequest{
		Host:             "127.0.0.1",
		Port:             uint32(rPort),
		Recover:          true,
		Forward:          true,
		SetAutoForward:   true,
		AutoForwardValue: true,
		EnableHttp2:      true,
	})

	var wg sync.WaitGroup
	wg.Add(1)
	for {
		rsp, err := stream.Recv()
		if err != nil {
			break
		}
		if strings.Contains(spew.Sdump(rsp), `starting mitm server`) && !started {
			println("----------------------")
			println("----------------------")
			println("----------------------")
			println("----------------------")
			println("----------------------")
			println("----------------------")
			println("----------------------")
			println("----------------------")
			println("----------------------")
			started = true
			go func() {
				defer func() {
					wg.Done()
					cancel()
					if err := recover(); err != nil {
						utils.PrintCurrentGoroutineRuntimeStack()
					}
				}()
				var token = utils.RandStringBytes(100)
				var params = map[string]any{
					"packet": lowhttp.ReplaceHTTPPacketHeader([]byte(`GET / HTTP/1.1
Host: www.example.com

`+token), "Host", utils.HostPort(mockHost, mockPort)),
					"host":  mockHost,
					"port":  mockPort,
					"proxy": proxy,
					"token": token,
				}
				spew.Dump(params)
				_, err := yak.NewScriptEngine(10).ExecuteEx(`
log.info("Start to send packet echo")
packet := getParam("packet")
host, port = getParam("host"), getParam("port")
dump(host, port, packet)
rsp, req = poc.HTTP(string(packet), poc.proxy(getParam("proxy")), poc.host(host), poc.port(port))~
dump(rsp)
dump(req)
if rsp.Contains(getParam("token")) {
		println("success")	
}else{
	dump(rsp)
	die("not pass!")
}
`, params)
				if err != nil {
					panic(err)
				}
				echoTested = true

				var tokenRaw, _ = utils.GzipCompress([]byte(token))
				params["packet"] = lowhttp.ReplaceHTTPPacketBody(utils.InterfaceToBytes(params["packet"]), tokenRaw, false)
				params["packet"] = lowhttp.ReplaceHTTPPacketHeader(utils.InterfaceToBytes(params["packet"]), "Content-Encoding", "gzip")
				_, err = yak.NewScriptEngine(10).ExecuteEx(`
log.info("Start to send packet echo")
packet := getParam("packet")
host, port = getParam("host"), getParam("port")
dump(host, port, packet)
rsp, req = poc.HTTP(string(packet), poc.proxy(getParam("proxy")), poc.host(host), poc.port(port))~
dump(rsp)
dump(req)
if rsp.Contains(getParam("token")) {
		println("success")	
}else{
	dump(rsp)
	die("not pass!")
}
`, params)
				if err != nil {
					panic(err)
				}
				gzipAutoDecode = true

				tokenRaw, _ = utils.GzipCompress([]byte(token))
				params["packet"] = lowhttp.ReplaceHTTPPacketBody(utils.InterfaceToBytes(params["packet"]), tokenRaw, false)
				params["packet"] = lowhttp.ReplaceHTTPPacketHeader(utils.InterfaceToBytes(params["packet"]), "Content-Encoding", "gzip")
				params["packet"] = lowhttp.HTTPPacketForceChunked(utils.InterfaceToBytes(params["packet"]))

				_, err = yak.NewScriptEngine(10).ExecuteEx(`
log.info("Start to send packet echo")
packet := getParam("packet")
host, port = getParam("host"), getParam("port")
dump(host, port, packet)
rsp, req = poc.HTTP(string(packet), poc.proxy(getParam("proxy")), poc.host(host), poc.port(port))~
dump(rsp)
if rsp.Contains(getParam("token")) {
		println("chunk + gzip auto decode success")	
}else{
	dump(rsp)
	die("not pass!")
}
`, params)
				if err != nil {
					panic(err)
				}
				chunkDecode = true

				tokenRaw = []byte(token)
				params["h2packet"] = lowhttp.ReplaceHTTPPacketBody([]byte(`GET /mitm/test/h2/token/`+token+` HTTP/2.0
Host: `+h2Addr+`
D: 1
`), tokenRaw, false)
				params["h2host"] = h2Host
				params["h2port"] = h2Port

				_, err = yak.NewScriptEngine(10).ExecuteEx(`
log.info("Start to send packet echo")
packet := getParam("h2packet")
println("-------------------------------------------------------------------------------------")
println("-------------------------------------------------------------------------------------")
println("-------------------------------------------------------------------------------------")
println("-------------------------------------------------------------------------------------")
println("-------------------------------------------------------------------------------------")
println("-------------------------------------------------------------------------------------")
dump(packet)
rsp, req = poc.HTTP(string(packet), poc.proxy(getParam("proxy")), poc.http2(true), poc.https(true))~
dump(rsp)
if rsp.Contains(getParam("token")) {
		println("h2 auto decode success")	
}else{
	dump(rsp)
	die("not pass!")
}
`, params)
				if err != nil {
					panic(err)
				}

				_, flows, err := yakit.QueryHTTPFlow(consts.GetGormProjectDatabase(), &ypb.QueryHTTPFlowRequest{
					SearchURL: "/mitm/test/h2/token/" + token,
				})
				if err != nil {
					panic(err)
				}
				spew.Dump(flows)
				if len(flows) > 0 {
					h2Test = true
				}
			}()
		}
		spew.Dump(rsp)
	}

	if !started {
		panic("MITM NOT STARTED!")
	}

	if !passthroughTested {
		panic("MITM PASSTHROUGH TEST FAILED")
	}

	if !echoTested {
		panic("MITM ECHO TEST FAILED")
	}

	if !gzipAutoDecode {
		panic("GZIP AUTO DECODE FAILED")
	}

	if !chunkDecode {
		panic("CHUNK DECODE FAILED")
	}

	if !h2Test {
		panic("H2 TEST FAILED")
	}
}
