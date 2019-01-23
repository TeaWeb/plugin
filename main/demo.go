package main

import (
	"github.com/TeaWeb/plugin/loader"
	"github.com/TeaWeb/plugin/plugins"
)

func main() {
	demoPlugin := plugins.NewPlugin()
	demoPlugin.Name = "Demo Plugin"
	demoPlugin.Code = "com.example.demo"
	demoPlugin.Developer = "Liu xiangchao"
	demoPlugin.Version = "1.0.0"
	demoPlugin.Date = "2018-10-15"
	demoPlugin.Site = "https://github.com/TeaWeb/build"
	demoPlugin.Description = "这是一个Demo插件"

	// HTTP请求筛选
	/**demoPlugin.OnRequest(func(request *http.Request) bool {
		log.Println("[demo]request:", request.URL.String())
		request.Header.Set("hello", "world")
		return true
	})**/

	// HTTP响应筛选
	/**demoPlugin.OnResponse(func(response *http.Response) bool {
		response.Header.Set("hello", "world")
		response.Header.Set("from", "demo")
		return true
	})**/

	loader.Start(demoPlugin)
}
