package teainterfaces

type PluginInterface interface {
	Name() string
	Code() string
	Site() string // 网站
	Version() string
	Date() string // 发布日期
	Developer() string
	Description() string

	Widgets() []interface{}
	OnLoad() error
	OnReload() error
	OnStart() error
	OnStop() error
	OnUnload() error
}
