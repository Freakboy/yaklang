package lowhttp

import (
	"strconv"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestExtractURLFromHTTPRequest(t *testing.T) {
	const packet = `GET / HTTP/1.1
Host: asdfasd:123
 Cookie: 123
  d: 1
`
	u, err := ExtractURLFromHTTPRequestRaw([]byte(packet), false)
	if err != nil {
		panic(err)
	}
	spew.Dump(u.String())
	var a = FixHTTPRequestOut([]byte(packet))
	if !strings.Contains(string(a), "\r\n Cookie: 123\r\n  d: 1\r\n") {
		panic(1)
	}
}

func TestParseStringToHttpRequest2(t *testing.T) {
	req, err := ParseStringToHttpRequest(`
GET / HTTP/1.1
Host: www.baidu.com

teadfasdfasd
`)
	if err != nil {
		t.FailNow()
		return
	}
	_ = req
}

func TestSplitHTTPHeader(t *testing.T) {
	key, value := SplitHTTPHeader("abc")
	if !(key == "abc" && value == "") {
		panic("111")
	}

	key, value = SplitHTTPHeader("abc:111")
	if !(key == "abc" && value == "111") {
		panic("111")
	}

	key, value = SplitHTTPHeader("abc: 111")
	if !(key == "abc" && value == "111") {
		panic("111")
	}

	key, value = SplitHTTPHeader("abc: 111\r\n")
	if !(key == "abc" && value == "111") {
		panic("111")
	}

	key, value = SplitHTTPHeader("Abc: 111\r\n")
	if !(key == "Abc" && value == "111") {
		panic("111")
	}

	key, value = SplitHTTPHeader("Abc: 1::11\r\n")
	if !(key == "Abc" && value == "1::11") {
		panic("111")
	}
}

func TestParseStringToHttpRequest(t *testing.T) {
	test := assert.New(t)

	req, err := ParseStringToHttpRequest(`
GET / HTTP/1.1
Host: www.baidu.com
Connection: close
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

`)
	if err != nil {
		test.FailNow(err.Error())
	}

	u, err := ExtractURLFromHTTPRequest(req, true)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	_ = u
}

func TestConvertContentToChunked(t *testing.T) {
	raws := fixInvalidHTTPHeaders([]byte(`
GET / HTTP/1.1
Host: www.baidu.com
Content-Length: 12

123123123123
`))
	println(string(raws))
}

func TestGetRedirectFromHTTPResponse(t *testing.T) {
	target := GetRedirectFromHTTPResponse([]byte(`HTTP/1.1 300 ...
Location: /target`), false)
	println(target)
	if target != "/target" {
		t.FailNow()
		return
	}
}

func TestRemoveZeroContentLengthHTTPHeader(t *testing.T) {
	target := RemoveZeroContentLengthHTTPHeader([]byte(`GET / HTTP/1.1
Host: www.baidu.com
Content-Length: 0

`))
	println(string(target))
	println(strconv.Quote(string(target)))
}

func TestFixHTTPResponse(t *testing.T) {
	rap, _, err := FixHTTPResponse([]byte(`HTTP/1.1 200 OK
Connection: close
Bdpagetype: 3
Bdqid: 0x9efbfb790011d570
Cache-Control: private
Ckpacknum: 2
Ckrndstr: 90011d570
Content-Encoding: gzip
Content-Type: text/html;charset=utf-8
Date: Sat, 27 Nov 2021 04:20:29 GMT
P3p: CP=" OTI DSP COR IVA OUR IND COM "
Server: BWS/1.1
Set-Cookie: BDRCVFR[S4-dAuiWMmn]=I67x6TjHwwYf0; path=/; domain=.baidu.com
Set-Cookie: delPer=0; path=/; domain=.baidu.com
Set-Cookie: BD_CK_SAM=1;path=/
Set-Cookie: PSINO=2; domain=.baidu.com; path=/
Set-Cookie: BDSVRTM=12; path=/
Set-Cookie: H_PS_PSSID=34445_35104_35239_34584_34517_35245_34606_35320_26350_35209_35312_35145; path=/; domain=.baidu.com
Strict-Transport-Security: max-age=172800
Traceid: 1637986829039139149811456026574257771888
Vary: Accept-Encoding
X-Frame-Options: sameorigin
X-Ua-Compatible: IE=Edge,chrome=1
Content-Length: 12

aaaaaaaaaaaa` + "\r\n\r\n"))
	if err != nil {
		return
	}
	println(string(rap))
}

func TestParseResponseLine(t *testing.T) {
	testcases := []struct {
		line          string
		proto, status string
		code          int
	}{
		{
			line:   "HTTP/1.1 200 OK",
			proto:  "HTTP/1.1",
			code:   200,
			status: "OK",
		},
		{
			line:   "HTTP/1.1 200",
			proto:  "HTTP/1.1",
			code:   200,
			status: "",
		},
		{
			line:   "HTTP/1.1 301 Moved Permanently",
			proto:  "HTTP/1.1",
			code:   301,
			status: "Moved Permanently",
		},
	}

	for _, testcase := range testcases {
		proto, code, status, _ := parseResponseLine(testcase.line)
		if proto != testcase.proto {
			t.Fatalf("parseResponseLine error: %s(got) != %s(want)", proto, testcase.proto)
		}
		if code != testcase.code {
			t.Fatalf("parseResponseLine error: %d(got) != %d(want)", code, testcase.code)
		}
		if status != testcase.status {
			t.Fatalf("parseResponseLine error: %s(got) != %s(want)", status, testcase.status)
		}

	}
}
