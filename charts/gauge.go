package charts

type GaugeChart struct {
	Chart

	value float64
	label string
	min   float64
	max   float64
	unit  string
}

func NewGaugeChart() *GaugeChart {
	p := &GaugeChart{}
	p.SetType("gauge")
	return p
}

func (this *GaugeChart) Value() float64 {
	return this.value
}

func (this *GaugeChart) SetValue(value float64) {
	this.value = value
}

func (this *GaugeChart) Label() string {
	return this.label
}

func (this *GaugeChart) SetLabel(label string) {
	this.label = label
}

func (this *GaugeChart) Min() float64 {
	return this.min
}

func (this *GaugeChart) SetMin(min float64) {
	this.min = min
}

func (this *GaugeChart) Max() float64 {
	return this.max
}

func (this *GaugeChart) SetMax(max float64) {
	this.max = max
}

func (this *GaugeChart) Unit() string {
	return this.unit
}

func (this *GaugeChart) SetUnit(unit string) {
	this.unit = unit
}
