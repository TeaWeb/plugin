package charts

type PieChart struct {
	Chart

	Values []interface{}
	Labels []string
}

func NewPieChart() *PieChart {
	p := &PieChart{
		Values: []interface{}{},
		Labels: []string{},
	}
	p.ChartType = "pie"
	return p
}
