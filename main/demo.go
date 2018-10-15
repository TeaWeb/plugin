package main

import (
	"github.com/TeaWeb/plugin/charts"
	"github.com/TeaWeb/plugin/plugins"
	"github.com/TeaWeb/plugin/teainterfaces"
)

func New() teainterfaces.PluginInterface {
	return &DemoPlugin{}
}

type DemoPlugin struct {
	plugins.Plugin
}

func (this *DemoPlugin) Name() string {
	return "Demo Plugin"
}

func (this *DemoPlugin) Code() string {
	return "com.example.demo"
}

func (this *DemoPlugin) Developer() string {
	return "Liu xiangchao"
}

func (this *DemoPlugin) Version() string {
	return "1.0.0"
}

func (this *DemoPlugin) Date() string {
	return "2018-10-15"
}

func (this *DemoPlugin) Site() string {
	return "https://github.com/TeaWeb/build"
}

func (this *DemoPlugin) Description() string {
	return "这是一个Demo插件"
}

func (this *DemoPlugin) OnLoad() error {
	widget := plugins.NewWidget()
	widget.SetName("Demo Chart")

	this.AddWidget(widget)

	// 添加一个进度条
	progressBar := charts.NewProgressBar()
	progressBar.SetValue(30)
	progressBar.SetName("ProgressBar")
	progressBar.SetDetail("ProgressBar Detail")
	widget.AddChart(progressBar)

	// 刷新Widget时的操作
	widget.SetOnReload(func() error {
		newValue := progressBar.Value() + 1
		if newValue > 100 {
			newValue = 100
		}
		progressBar.SetValue(newValue)
		return nil
	})

	// 添加一个表格
	table := charts.NewTable()
	table.SetName("Table")
	table.SetDetail("Table Detail")
	table.AddRow("Col1", "Col2")
	table.AddRow("Col3", "Col4")
	table.SetWidth(30, 70)
	widget.AddChart(table)

	// 添加一个仪表盘
	gauge := charts.NewGaugeChart()
	gauge.SetName("Gauge")
	gauge.SetDetail("MB")
	gauge.SetDetail("Gauge Detail")
	gauge.SetMax(20)
	gauge.SetValue(15)
	widget.AddChart(gauge)

	// 添加一个饼图
	pie := charts.NewPieChart()
	pie.SetName("Pie")
	pie.SetDetail("Pie Detail")
	pie.SetValues([]interface{}{1, 2, 3})
	pie.SetLabels([]string{"A", "B", "C"})
	widget.AddChart(pie)

	// 添加一个线图
	line := charts.NewLine()
	line.SetName("Line 1")
	line.SetValues([]interface{}{1, 2, 3, 2, 1})
	line.SetFilled(true)
	line.SetColor(teainterfaces.ColorBlue)

	lineChart := charts.NewLineChart()
	lineChart.SetName("Line")
	lineChart.SetDetail("Line Detail")
	lineChart.AddLine(line)
	lineChart.SetLabels([]string{"A", "B", "C", "D", "E"})
	widget.AddChart(lineChart)

	return nil
}
