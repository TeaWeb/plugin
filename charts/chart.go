package charts

type Chart struct {
	id        string
	chartType string
	name      string
	detail    string
}

func (this *Chart) Id() string {
	return this.id
}

func (this *Chart) SetId(id string) {
	this.id = id
}

func (this *Chart) Type() string {
	return this.chartType
}

func (this *Chart) SetType(chartType string) {
	this.chartType = chartType
}

func (this *Chart) Name() string {
	return this.name
}

func (this *Chart) SetName(name string) {
	this.name = name
}

func (this *Chart) Detail() string {
	return this.detail
}

func (this *Chart) SetDetail(detail string) {
	this.detail = detail
}
