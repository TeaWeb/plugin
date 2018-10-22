package charts

type GaugeChart struct {
	Chart

	Value float64
	Label string
	Min   float64
	Max   float64
	Unit  string
}

func NewGaugeChart() *GaugeChart {
	p := &GaugeChart{}
	p.ChartType = "gauge"
	return p
}
