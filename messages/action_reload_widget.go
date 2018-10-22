package messages

type ReloadWidgetAction struct {
	Action

	WidgetId string
}

func (this *ReloadWidgetAction) Name() string {
	return "ReloadWidget"
}
