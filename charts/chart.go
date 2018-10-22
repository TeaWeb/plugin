package charts

type ChartInterface interface {
	ChartId() string
	SetId(id string)

	OnReload(f func())
	Reload()

	NotifyChange()
	IsChanged() bool
}

type Chart struct {
	Id        string
	ChartType string
	Name      string
	Detail    string

	isChanged    bool
	onReloadFunc func()
}

func (this *Chart) ChartId() string {
	return this.Id
}

func (this *Chart) SetId(id string) {
	this.Id = id
}

func (this *Chart) OnReload(f func()) {
	this.onReloadFunc = f
}

func (this *Chart) Reload() {
	if this.onReloadFunc != nil {
		this.onReloadFunc()
	}
}

func (this *Chart) NotifyChange() {
	this.isChanged = true
}

func (this *Chart) IsChanged() bool {
	if !this.isChanged {
		return false
	}
	this.isChanged = false
	return true
}
