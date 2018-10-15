package plugins

import (
	"github.com/TeaWeb/plugin/teainterfaces"
)

func NewWidget() *Widget {
	return &Widget{
		group:     teainterfaces.WidgetGroupService,
		dashboard: true,
		charts:    []interface{}{},
	}
}

type Widget struct {
	name      string
	icon      []byte
	title     string
	url       string
	moreURL   string
	topBar    bool
	menuBar   bool
	helperBar bool
	dashboard bool
	group     teainterfaces.WidgetGroup
	charts    []interface{}

	onReloadFunc func() error
}

func (this *Widget) Name() string {
	return this.name
}

func (this *Widget) SetName(name string) {
	this.name = name
}

func (this *Widget) Icon() []byte {
	return this.icon
}

func (this *Widget) SetIcon(icon []byte) {
	this.icon = icon
}

func (this *Widget) Title() string {
	return this.title
}

func (this *Widget) SetTitle(title string) {
	this.title = title
}

func (this *Widget) URL() string {
	return this.url
}

func (this *Widget) SetURL(url string) {
	this.url = url
}

func (this *Widget) MoreURL() string {
	return this.moreURL
}

func (this *Widget) SetMoreURL(moreURL string) {
	this.moreURL = moreURL
}

func (this *Widget) TopBar() bool {
	return this.topBar
}

func (this *Widget) ShowInTopBar(b bool) {
	this.topBar = b
}

func (this *Widget) MenuBar() bool {
	return this.menuBar
}

func (this *Widget) ShowInMenuBar(b bool) {
	this.menuBar = b
}

func (this *Widget) HelperBar() bool {
	return this.helperBar
}

func (this *Widget) ShowInHelperBar(b bool) {
	this.helperBar = b
}

func (this *Widget) Dashboard() bool {
	return this.dashboard
}

func (this *Widget) ShowInDashboard(b bool) {
	this.dashboard = b
}

func (this *Widget) Group() teainterfaces.WidgetGroup {
	return this.group
}

func (this *Widget) SetGroup(group teainterfaces.WidgetGroup) {
	this.group = group
}

func (this *Widget) Charts() []interface{} {
	return this.charts
}

func (this *Widget) AddChart(chart teainterfaces.ChartInterface) {
	this.charts = append(this.charts, chart)
}

func (this *Widget) OnReload() error {
	if this.onReloadFunc != nil {
		return this.onReloadFunc()
	}
	return nil
}

func (this *Widget) SetOnReload(f func() error) {
	this.onReloadFunc = f
}
