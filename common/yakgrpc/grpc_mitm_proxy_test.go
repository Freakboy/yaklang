package yakgrpc

import (
	"context"
	"fmt"
	"github.com/yaklang/yaklang/common/crep"
	"github.com/yaklang/yaklang/common/utils"
	"github.com/yaklang/yaklang/common/yak"
	"github.com/yaklang/yaklang/common/yakgrpc/ypb"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TemplateTestGRPCMUSTPASS_MITM_Proxy_Template(t *testing.T) {
	ctx := utils.TimeoutContextSeconds(10)
	port := utils.GetRandomAvailableTCPPort()
	server, err := crep.NewMITMServer()
	if err != nil {
		panic(err)
	}
	addr := utils.HostPort("127.0.0.1", port)
	go func() {
		server.Serve(ctx, addr)
	}()
	if utils.WaitConnect(addr, 10) != nil {
		panic("wait connect timeout")
	}

	mitmPort := utils.GetRandomAvailableTCPPort()
	client, err := NewLocalClient()
	if err != nil {
		panic(err)
	}
	stream, err := client.MITM(ctx)
	if err != nil {
		panic(err)
	}
	stream.Send(&ypb.MITMRequest{
		Host:            "127.0.0.1",
		Port:            uint32(mitmPort),
		DownstreamProxy: "http://" + utils.HostPort("127.0.0.1", port),
	})
	for {
		data, err := stream.Recv()
		if err != nil {
			break
		}
		if data.GetMessage().GetIsMessage() {
			var msg = string(data.GetMessage().GetMessage())
			fmt.Println(msg)
			if strings.Contains(msg, "starting mitm server") {
				// do sth
			}
		}
	}
}

func TestGRPCMUSTPASS_MITM_Proxy(t *testing.T) {
	var (
		networkIsPassed  bool
		downstreamPassed bool
		token            = utils.RandNumberStringBytes(10)
	)
	mockHost, mockPort := utils.DebugMockHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("u") == token {
			networkIsPassed = true
		}
		writer.Write([]byte("Hello Token"))
	})
	var mockUrl = "http://" + utils.HostPort(mockHost, mockPort)

	ctx := utils.TimeoutContextSeconds(10)
	port := utils.GetRandomAvailableTCPPort()
	server, err := crep.NewMITMServer(crep.MITM_SetHTTPRequestHijack(func(https bool, req *http.Request) *http.Request {
		if req.URL.Query().Get("u") == token {
			downstreamPassed = true
		}
		return req
	}))
	if err != nil {
		panic(err)
	}
	addr := utils.HostPort("127.0.0.1", port)
	go func() {
		server.Serve(ctx, addr)
	}()
	if utils.WaitConnect(addr, 10) != nil {
		panic("wait connect timeout")
	}

	mitmPort := utils.GetRandomAvailableTCPPort()
	client, err := NewLocalClient()
	if err != nil {
		panic(err)
	}
	stream, err := client.MITM(ctx)
	if err != nil {
		panic(err)
	}
	stream.Send(&ypb.MITMRequest{
		Host:            "127.0.0.1",
		Port:            uint32(mitmPort),
		DownstreamProxy: "http://" + utils.HostPort("127.0.0.1", port),
	})
	for {
		data, err := stream.Recv()
		if err != nil {
			break
		}
		if data.GetMessage().GetIsMessage() {
			var msg = string(data.GetMessage().GetMessage())
			fmt.Println(msg)
			if strings.Contains(msg, "starting mitm server") {
				if _, err := yak.Execute(
					`
dump(mockUrl, token, mitmProxy)
poc.Get(mockUrl, poc.proxy(mitmProxy), poc.replaceQueryParam("u", token))~`,
					map[string]any{
						"mockUrl":   mockUrl,
						"mitmProxy": "http://" + utils.HostPort("127.0.0.1", mitmPort),
						"token":     token,
					}); err != nil {
					t.Errorf("execute script failed: %v", err)
					t.FailNow()
				}
			}
		}
	}

	if !downstreamPassed {
		t.Fatalf("Downstream proxy not passed")
	}

	if !networkIsPassed {
		t.Fatalf("Network not passed")
	}
	t.Log("PASS")
}

func TestGRPCMUSTPASS_MITM_Proxy_ApplyToPlugin(t *testing.T) {
	var (
		networkIsPassed  bool
		downstreamPassed bool
		token            = utils.RandNumberStringBytes(10)

		pluginNetworkIsSPassed bool
		pluginDownstreamPassed bool
		tokenForPlugin         = utils.RandNumberStringBytes(30)
	)
	mockHost, mockPort := utils.DebugMockHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("u") == token {
			networkIsPassed = true
		}

		if request.URL.Query().Get("u") == tokenForPlugin && request.Method == "POST" {
			pluginNetworkIsSPassed = true
		}
		writer.Write([]byte("Hello Token"))
	})
	var mockUrl = "http://" + utils.HostPort(mockHost, mockPort)
	_ = mockUrl

	ctx, cancel := context.WithCancel(utils.TimeoutContextSeconds(10))
	port := utils.GetRandomAvailableTCPPort()
	server, err := crep.NewMITMServer(crep.MITM_SetHTTPRequestHijack(func(https bool, req *http.Request) *http.Request {
		if req.URL.Query().Get("u") == token {
			downstreamPassed = true
		}

		if req.URL.Query().Get("u") == tokenForPlugin && req.Method == "POST" {
			pluginDownstreamPassed = true
		}
		return req
	}))
	if err != nil {
		panic(err)
	}
	addr := utils.HostPort("127.0.0.1", port)
	go func() {
		server.Serve(ctx, addr)
	}()
	if utils.WaitConnect(addr, 10) != nil {
		panic("wait connect timeout")
	}

	mitmPort := utils.GetRandomAvailableTCPPort()
	client, err := NewLocalClient()
	if err != nil {
		panic(err)
	}
	stream, err := client.MITM(ctx)
	if err != nil {
		panic(err)
	}
	stream.Send(&ypb.MITMRequest{
		Host:            "127.0.0.1",
		Port:            uint32(mitmPort),
		DownstreamProxy: "http://" + utils.HostPort("127.0.0.1", port),
	})
	for {
		data, err := stream.Recv()
		if err != nil {
			break
		}
		if data.GetMessage().GetIsMessage() {
			var msg = string(data.GetMessage().GetMessage())
			fmt.Println(msg)
			if strings.Contains(msg, `Initializing HotPatched MITM HOOKS`) {
				if _, err := yak.Execute(
					`
poc.Get(mockUrl, poc.proxy(mitmProxy), poc.replaceQueryParam("u", token))~`,
					map[string]any{
						"mockUrl":   mockUrl,
						"mitmProxy": "http://" + utils.HostPort("127.0.0.1", mitmPort),
						"token":     token,
					}); err != nil {
					t.Errorf("execute script failed: %v", err)
					cancel()
					t.FailNow()
				}
			}
			if strings.Contains(msg, "starting mitm server") {
				stream.Send(&ypb.MITMRequest{
					SetYakScript: true,
					YakScriptContent: `mirrorNewWebsite = (https, url, req, rsp, body) => {
	println("DO: %v" % url)
	if str.HasPrefix(req, "CONNECT ") {return}
	poc.Post(url, poc.replaceQueryParam("u", ` + strconv.Quote(tokenForPlugin) + `))~
}`,
				})

			}
		}
	}

	if !downstreamPassed {
		t.Fatalf("Downstream proxy not passed")
	}

	if !networkIsPassed {
		t.Fatalf("Network not passed")
	}

	if !pluginNetworkIsSPassed {
		t.Fatalf("Plugin Network not passed")
	} else {
		t.Log("Plugin Network passed")
	}

	if !pluginDownstreamPassed {
		t.Fatalf("Plugin Downstream not passed")
	}
	t.Log("PASS")
}
