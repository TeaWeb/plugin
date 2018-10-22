package main

import (
	"github.com/TeaWeb/plugin/charts"
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

	// 添加widget
	addWidget(demoPlugin)

	// 请求筛选
	/**demoPlugin.OnRequest(func(request *http.Request) bool {
		log.Println("[demo]request:", request.URL.String())
		request.Header.Set("hello", "world")
		return true
	})**/

	loader.Start(demoPlugin)
}

// 添加Widget
func addWidget(plugin *plugins.Plugin) {
	widget := plugins.NewWidget()
	widget.Name = "Demo Chart"

	plugin.AddWidget(widget)

	// 添加一个进度条
	progressBar := charts.NewProgressBar()
	progressBar.Value = 30
	progressBar.Name = "ProgressBar"
	progressBar.Detail = "ProgressBar Detail"
	widget.AddChart(progressBar)

	// 刷新Widget时的操作
	widget.OnReload(func() {
		newValue := progressBar.Value + 1
		if newValue > 100 {
			newValue = 100
		}
		progressBar.Value = newValue
		progressBar.NotifyChange()
	})

	// 添加一个表格
	table := charts.NewTable()
	table.Name = "Table"
	table.Detail = "Table Detail"
	table.AddRow("Col1", "Col2")
	table.AddRow("Col3", "Col4")
	table.SetWidth(30, 70)
	widget.AddChart(table)

	// 添加一个仪表盘
	gauge := charts.NewGaugeChart()
	gauge.Name = "Gauge"
	gauge.Unit = "MB"
	gauge.Detail = "Gauge Detail"
	gauge.Max = 20
	gauge.Value = 15
	widget.AddChart(gauge)

	// 添加一个饼图
	pie := charts.NewPieChart()
	pie.Name = "Pie"
	pie.Detail = "Pie Detail"
	pie.Values = []interface{}{1, 2, 3}
	pie.Labels = []string{"A", "B", "C"}
	widget.AddChart(pie)

	// 添加一个线图
	line := charts.NewLine()
	line.Name = "Line 1"
	line.Values = []interface{}{1, 2, 3, 2, 1}
	line.Filled = true
	line.Color = charts.ColorBlue

	lineChart := charts.NewLineChart()
	lineChart.Name = "Line"
	lineChart.Detail = "Line Detail"
	lineChart.AddLine(line)
	lineChart.Labels = []string{"A", "B", "C", "D", "E"}
	widget.AddChart(lineChart)
}
