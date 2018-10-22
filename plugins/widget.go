package plugins

import (
	"github.com/TeaWeb/plugin/charts"
)

func NewWidget() *Widget {
	return &Widget{
		Group:     WidgetGroupService,
		Dashboard: true,
		Charts:    []interface{}{},
	}
}

type Widget struct {
	Id      string
	Name    string
	Icon    []byte
	Title   string
	URL     string
	MoreURL string
	TopBar  bool

	MenuBar   bool
	HelperBar bool
	Dashboard bool

	Group  WidgetGroup
	Charts []interface{}

	onReloadFunc func()
}

func (this *Widget) AddChart(chart charts.ChartInterface) {
	if len(chart.ChartId()) == 0 {
		chart.SetId(RandString(16))
	}
	this.Charts = append(this.Charts, chart)
}

func (this *Widget) OnReload(f func()) {
	this.onReloadFunc = f
}

func (this *Widget) Reload() {
	if this.onReloadFunc != nil {
		this.onReloadFunc()
	}

	// 刷新charts
	for _, chart := range this.Charts {
		_, ok := chart.(charts.ChartInterface)
		if ok {
			chart.(charts.ChartInterface).Reload()
		}
	}
}
