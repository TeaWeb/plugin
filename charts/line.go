package charts

func NewLine() *Line {
	return &Line{}
}

type Line struct {
	Name      string
	Values    []interface{}
	Color     Color
	Filled    bool
	ShowItems bool
}

func NewLineChart() *LineChart {
	p := &LineChart{}
	p.ChartType = "line"
	p.Lines = []*Line{}
	return p
}

type LineChart struct {
	Chart

	Lines  []*Line
	Labels []string

	Max       float64
	XShowTick bool // X轴是否显示刻度

	YTickCount uint // Y轴刻度分隔数量
	YShowTick  bool // Y轴是否显示刻度
}

func (this *LineChart) AddLine(line *Line) {
	this.Lines = append(this.Lines, line)
}

func (this *LineChart) ResetLines() {
	this.Lines = []*Line{}
}
