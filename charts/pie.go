package charts

type PieChart struct {
	Chart

	values []interface{}
	labels []string
}

func NewPieChart() *PieChart {
	p := &PieChart{
		values: []interface{}{},
		labels: []string{},
	}
	p.SetType("pie")
	return p
}

func (this *PieChart) Values() []interface{} {
	return this.values
}

func (this *PieChart) SetValues(values []interface{}) {
	this.values = values
}

func (this *PieChart) Labels() []string {
	return this.labels
}

func (this *PieChart) SetLabels(labels []string) {
	this.labels = labels
}
