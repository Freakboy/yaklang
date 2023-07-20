package vulinbox

import (
	"encoding/json"
	"fmt"
	"github.com/yaklang/yaklang/common/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (s *VulinServer) registerSSRF() {

	ssrfGroup := s.router.PathPrefix("/ssrf").Subrouter()

	ssrfGroup.HandleFunc("/json-in-get", func(writer http.ResponseWriter, request *http.Request) {
		raw := request.URL.Query().Get("json")
		if raw == "" {
			writer.Write([]byte(`暂无数据！`))
			return
		}
		var m = make(map[string]interface{})
		err := json.Unmarshal([]byte(raw), &m)
		if err != nil {
			writer.Write([]byte(`JSON Syntax Error: ` + err.Error()))
			return
		}

		ref, ok := m["ref"]
		if !ok {
			writer.Write([]byte(`no ref in json found!`))
			return
		}

		var u = fmt.Sprint(ref)
		c := utils.NewDefaultHTTPClient()
		c.Timeout = 5 * time.Second
		rsp, err := c.Get(u)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		rawResponse, err := utils.HttpDumpWithBody(rsp, true)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(rawResponse)
	}).Name("SSRF JSON Body SSRF")
	ssrfGroup.HandleFunc("/in-get", func(writer http.ResponseWriter, request *http.Request) {
		ref := request.URL.Query().Get("url")
		var u = fmt.Sprint(ref)
		c := utils.NewDefaultHTTPClient()
		c.Timeout = 5 * time.Second
		rsp, err := c.Get(u)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		rawResponse, err := utils.HttpDumpWithBody(rsp, true)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(rawResponse)
	}).Queries("url", "").Name("SSRF GET 中 URL 参数")
	ssrfGroup.HandleFunc("/in-post", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			writer.Header().Set("Content-Type", "text/html; charset=utf8")
			writer.Write([]byte(`
<form action="/ssrf-in-post" method="post">
	<label for="name">姓名:</label>
	<input type="text" id="name" name="name" ><br><br>

	<label for="email">邮箱:</label>
	<input type="email" id="email" name="email" ><br><br>

	<label for="age">年龄:</label>
	<input type="number" id="age" name="age" min="2" max="120" ><br><br>

	
	<label for="url">URL</label>
	<input id='url' name="url"><br><br>

	<label for="gender">性别:</label>
	<select id="gender" name="gender" >
		<option value="">请选择</option>
		<option value="male">男</option>
		<option value="female">女</option>
		<option value="other">其他</option>
	</select><br><br>

	<label for="message">留言:</label>
	<textarea id="message" name="message" rows="5" ></textarea><br><br>

	<input type="submit" value="提交">
</form>
`))
			return
		}
		raw, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		values, err := url.ParseQuery(string(raw))
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		var u = fmt.Sprint(values.Get("url"))
		c := utils.NewDefaultHTTPClient()
		c.Timeout = 10 * time.Second
		rsp, err := c.Get(u)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		rawResponse, err := utils.HttpDumpWithBody(rsp, true)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(rawResponse)
	}).Name("SSRF POST 中 URL 参数")
	ssrfGroup.HandleFunc("/in-post-multipart", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			writer.Header().Set("Content-Type", "text/html; charset=utf8")
			writer.Write([]byte(`
<form action="/ssrf-in-post" method="post" enctype="multipart/form-data">
	<label for="name">姓名:</label>
	<input type="text" id="name" name="name" ><br><br>

	<label for="email">邮箱:</label>
	<input type="email" id="email" name="email" ><br><br>

	<label for="age">年龄:</label>
	<input type="number" id="age" name="age" min="2" max="120" ><br><br>

	
	<label for="url">URL</label>
	<input id='url' name="url"><br><br>

	<label for="gender">性别:</label>
	<select id="gender" name="gender" >
		<option value="">请选择</option>
		<option value="male">男</option>
		<option value="female">女</option>
		<option value="other">其他</option>
	</select><br><br>

	<label for="message">留言:</label>
	<textarea id="message" name="message" rows="5" ></textarea><br><br>

	<input type="submit" value="提交">
</form>
`))
			return
		}
	})
	s.router.HandleFunc("/redirect/main", func(writer http.ResponseWriter, request *http.Request) {
		DefaultRender(`<h1>Hello, Welcome to Vulinbox!</h1>`, writer, request)
	})
	s.router.HandleFunc("/redirect/basic", func(writer http.ResponseWriter, request *http.Request) {
		var u = LoadFromGetParams(request, "destUrl")
		if strings.Contains(u, `redirect/basic`) {
			DefaultRender("<p>forbidden to "+strconv.Quote(u)+"</p>", writer, request)
			return
		}
		writer.Header().Set("Location", u)
		writer.WriteHeader(302)
	})
	s.router.HandleFunc("/redirect/redirect-hell", func(writer http.ResponseWriter, request *http.Request) {
		var u = LoadFromGetParams(request, "destUrl")
		writer.Header().Set("Location", u)
		writer.WriteHeader(302)
	})
	s.router.HandleFunc("/redirect/js/basic", func(writer http.ResponseWriter, request *http.Request) {
		var u = LoadFromGetParams(request, "redUrl")
		DefaultRender(`
<h2>Open Redirect With JS</h2>
<a href=`+strconv.Quote(u)+`>Click ME JUMP NOW (3s)</a>
<script>
	setTimeout(function() {

window.location.href = `+strconv.Quote(u)+`;

}, 3000)
</script>
`, writer, request)
	})
	s.router.HandleFunc("/redirect/js/basic1", func(writer http.ResponseWriter, request *http.Request) {
		var u = LoadFromGetParams(request, "redirect_to")
		DefaultRender(`
<h2>Open Redirect With JS</h2>
<a href=`+strconv.Quote(u)+`>Click ME JUMP NOW (3s)</a>
<script>
	setTimeout(function() {

window.location.replace(`+strconv.Quote(u)+`);

}, 3000)
</script>
`, writer, request)
	})
	s.router.HandleFunc("/redirect/js/basic2", func(writer http.ResponseWriter, request *http.Request) {
		var u = LoadFromGetParams(request, "redirect")
		DefaultRender(`
<h2>Open Redirect With JS</h2>
<a href=`+strconv.Quote(u)+`>Click ME JUMP NOW (3s)</a>
<script>
	setTimeout(function() {

window.location.assign(`+strconv.Quote(u)+`);

}, 3000)
</script>
`, writer, request)
	})
	s.router.HandleFunc("/redirect/meta/case1", func(writer http.ResponseWriter, request *http.Request) {
		var u = LoadFromGetParams(request, "redirect")
		DefaultRenderEx(true, `<!DOCTYPE html>
<html>
  <head>
    <title>Meta(5s) Refresh Example</title>
    <meta http-equiv="refresh" content="5;url={{ .url }}">
  </head>
</html>
`, writer, request, map[string]any{
			"url": strings.Trim(strconv.Quote(u), `"`),
		})
	})
	s.router.HandleFunc("/redirect/meta/case2", func(writer http.ResponseWriter, request *http.Request) {
		var u = LoadFromGetParams(request, "redirect")
		DefaultRenderEx(true, `<!DOCTYPE html>
<html>
  <head>
    <title>Meta Refresh Example</title>
    <meta http-equiv="refresh" content="0;url={{ .url }}">
  </head>
</html>
`, writer, request, map[string]any{
			"url": strings.Trim(strconv.Quote(u), `"`),
		})
	})
	}).Name("SSRF POST 中 URL 参数(Multipart)")
	ssrfGroup.HandleFunc("/in-json-body", func(writer http.ResponseWriter, request *http.Request) {

		return
	}).Name("SSRF JSON Body SSRF")
	ssrfGroup.HandleFunc("/json-in-post-param", func(writer http.ResponseWriter, request *http.Request) {

		return
	}).Name("SSRF POST参数是JSON（包含URL）的情况")
}
