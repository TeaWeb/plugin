package plugins

import (
	"github.com/TeaWeb/plugin/apps"
	"net/http"
)

func NewPlugin() *Plugin {
	return &Plugin{}
}

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

	onRequestFunc  func(request *http.Request) bool
	onResponseFunc func(response *http.Response, writer http.ResponseWriter) bool
}

func (this *Plugin) OnReload(f func()) {
	this.onReloadFunc = f
}

func (this *Plugin) Reload() {
	if this.onReloadFunc != nil {
		this.onReloadFunc()
	}
}

func (this *Plugin) OnStart(f func()) {
	this.onStartFunc = f
}

func (this *Plugin) Start() {
	if this.onStartFunc != nil {
		this.onStartFunc()
	}
}

func (this *Plugin) OnStop(f func()) {
	this.onStopFunc = f
}

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

// 添加App
func (this *Plugin) AddApp(app *apps.App) {
	if len(app.Id) == 0 {
		app.Id = RandString(16)
	}
	this.Apps = append(this.Apps, app)
}

// 过滤请求，如果返回false，则不会往下执行
func (this *Plugin) OnRequest(f func(request *http.Request) bool) {
	this.onRequestFunc = f

	if this.onRequestFunc != nil {
		this.HasRequestFilter = true
	}
}

func (this *Plugin) FilterRequest(request *http.Request) bool {
	if this.onRequestFunc != nil {
		return this.onRequestFunc(request)
	}
	return true
}

// 过滤响应，如果返回false，则不会往下执行
// TODO 需要实现
func (this *Plugin) OnResponse(f func(response *http.Response, writer http.ResponseWriter) bool) {
	this.onResponseFunc = f

	if this.onResponseFunc != nil {
		this.HasResponseFilter = true
	}
}
