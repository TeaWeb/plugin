package messages

func init() {
	RegisterAction(
		new(RegisterPluginAction),
		new(LoadAction),
		new(StartAction),
		new(StopAction),
		new(ReloadAction),
		new(ReloadAppAction),
		new(ReloadAppsAction),
		new(FilterRequestAction),
		new(FilterResponseAction),
	)
}
