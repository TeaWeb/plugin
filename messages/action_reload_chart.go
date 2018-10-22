package messages

type ReloadChartAction struct {
	Action

	Chart interface{}
}

func (this *ReloadChartAction) Name() string {
	return "ReloadChart"
}
