package plugins

type WidgetGroup = uint8

const (
	WidgetGroupSystem   = WidgetGroup(1) // 系统信息
	WidgetGroupService  = WidgetGroup(2) // 服务
	WidgetGroupRealTime = WidgetGroup(3) // 即时
)
