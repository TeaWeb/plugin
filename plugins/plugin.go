package plugins

import (
	"github.com/TeaWeb/plugin/apps"
	"net/http"
)

// 构造新的插件对象
func NewPlugin() *Plugin {
	return &Plugin{}
}

// 插件定义
type Plugin struct {
	Name        string
	Code        string
	Site        string
	Version     string
	Date        string
	Developer   string
	Description string

	Widgets []*Widget
	Apps    []*apps.App

	HasRequestFilter  bool
	HasResponseFilter bool

	onReloadFunc func()
	onStartFunc  func()
	onStopFunc   func()

	onReloadAppsFunc func()

	onRequestFunc  func(request *http.Request) bool
	onResponseFunc func(response *http.Response) bool
}

// 设置刷新时回调函数
func (this *Plugin) OnReload(f func()) {
	this.onReloadFunc = f
}

// 刷新插件
func (this *Plugin) Reload() {
	if this.onReloadFunc != nil {
		this.onReloadFunc()
	}
}

// 设置启动时回调函数
func (this *Plugin) OnStart(f func()) {
	this.onStartFunc = f
}

// 启动插件
func (this *Plugin) Start() {
	if this.onStartFunc != nil {
		this.onStartFunc()
	}
}

// 设置停止时回调函数
func (this *Plugin) OnStop(f func()) {
	this.onStopFunc = f
}

// 停止插件
func (this *Plugin) Stop() {
	if this.onStopFunc != nil {
		this.onStopFunc()
	}
}

// 添加Widget
func (this *Plugin) AddWidget(widget *Widget) {
	if len(widget.Id) == 0 {
		widget.Id = RandString(16)
	}
	this.Widgets = append(this.Widgets, widget)
}

// 根据ID获取Widget
func (this *Plugin) WidgetWithId(widgetId string) *Widget {
	for _, widget := range this.Widgets {
		if widget.Id == widgetId {
			return widget
		}
	}
	return nil
}

// 刷新apps时回调
func (this *Plugin) OnReloadApps(f func()) {
	this.onReloadAppsFunc = f
}

// 刷新apps
func (this *Plugin) ReloadApps() {
	if this.onReloadAppsFunc != nil {
		this.onReloadAppsFunc()
	}
}

// 添加App
func (this *Plugin) AddApp(app ... *apps.App) {
	for _, a := range app {
		if len(a.Id) == 0 {
			a.Id = RandString(16)
		}
		this.Apps = append(this.Apps, a)
	}

	if len(this.Apps) > 0 {
		this.ReloadApps()
	}
}

// 根据ID获取App
func (this *Plugin) AppWithId(appId string) *apps.App {
	for _, app := range this.Apps {
		if app.Id == appId {
			return app
		}
	}
	return nil
}

// 设置请求时的回调函数，如果返回false，则不会往下执行
func (this *Plugin) OnRequest(f func(request *http.Request) bool) {
	this.onRequestFunc = f

	if this.onRequestFunc != nil {
		this.HasRequestFilter = true
	}
}

// 调用请求时的回调函数
func (this *Plugin) FilterRequest(request *http.Request) bool {
	if this.onRequestFunc != nil {
		return this.onRequestFunc(request)
	}
	return true
}

// 设置发送响应时的回调函数，如果返回false，则不会往下执行
func (this *Plugin) OnResponse(f func(response *http.Response) bool) {
	this.onResponseFunc = f

	if this.onResponseFunc != nil {
		this.HasResponseFilter = true
	}
}

// 调用发送响应时的回调函数
func (this *Plugin) FilterResponse(response *http.Response) bool {
	if this.onResponseFunc != nil {
		return this.onResponseFunc(response)
	}
	return true
}
