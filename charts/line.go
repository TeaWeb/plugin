package charts

import "github.com/TeaWeb/plugin/teainterfaces"

func NewLine() *Line {
	return &Line{}
}

type Line struct {
	name      string
	values    []interface{}
	color     teainterfaces.Color
	filled    bool
	showItems bool
}

func (this *Line) Name() string {
	return this.name
}

func (this *Line) SetName(name string) {
	this.name = name
}

func (this *Line) Values() []interface{} {
	return this.values
}

func (this *Line) SetValues(values []interface{}) {
	this.values = values
}

func (this *Line) Color() teainterfaces.Color {
	return this.color
}

func (this *Line) SetColor(color teainterfaces.Color) {
	this.color = color
}

func (this *Line) Filled() bool {
	return this.filled
}

func (this *Line) SetFilled(filled bool) {
	this.filled = filled
}

func (this *Line) ShowItems() bool {
	return this.showItems
}

func (this *Line) SetShowItems(showItems bool) {
	this.showItems = showItems
}

func NewLineChart() *LineChart {
	p := &LineChart{}
	p.SetType("line")
	p.lines = []interface{}{}
	return p
}

type LineChart struct {
	Chart

	lines  []interface{}
	labels []string

	max       float64
	xShowTick bool // X轴是否显示刻度

	yTickCount uint // Y轴刻度分隔数量
	yShowTick  bool // Y轴是否显示刻度
}

func (this *LineChart) Lines() []interface{} {
	return this.lines
}

func (this *LineChart) Labels() []string {
	return this.labels
}

func (this *LineChart) SetLabels(labels []string) {
	this.labels = labels
}

func (this *LineChart) Max() float64 {
	return this.max
}

func (this *LineChart) SetMax(max float64) {
	this.max = max
}

func (this *LineChart) XShowTick() bool {
	return this.xShowTick
}

func (this *LineChart) SetXShowTick(b bool) {
	this.xShowTick = b
}

func (this *LineChart) SetYTickCount(count uint) {
	this.yTickCount = count
}

func (this *LineChart) YTickCount() uint {
	return this.yTickCount
}

func (this *LineChart) YShowTick() bool {
	return this.yShowTick
}

func (this *LineChart) SetYShowTick(b bool) {
	this.yShowTick = b
}

func (this *LineChart) AddLine(line *Line) {
	this.lines = append(this.lines, line)
}

func (this *LineChart) ResetLines() {
	this.lines = []interface{}{}
}
