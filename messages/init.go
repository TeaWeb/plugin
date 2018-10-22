package messages

func init() {
	RegisterAction(
		new(RegisterPluginAction),
		new(LoadAction),
		new(StartAction),
		new(StopAction),
		new(ReloadAction),
		new(ReloadWidgetAction),
		new(ReloadChartAction),
		new(FilterRequestAction),
		new(FilterResponseAction),
	)
}
