package plugins

import "github.com/TeaWeb/plugin/teainterfaces"

type PluginWidgets struct {
	widgets []interface{}
}

func (this *PluginWidgets) Widgets() []interface{} {
	return this.widgets
}

func (this *PluginWidgets) AddWidget(widget teainterfaces.WidgetInterface) {
	if this.widgets == nil {
		this.widgets = []interface{}{widget}
	} else {
		this.widgets = append(this.widgets, widget)
	}
}

type Plugin struct {
}

func (this *Plugin) Site() string {
	return ""
}

func (this *Plugin) Version() string {
	return "1.0.0"
}

func (this *Plugin) Date() string {
	return ""
}

func (this *Plugin) Developer() string {
	return "UNKNOWN"
}

func (this *Plugin) Description() string {
	return ""
}

func (this *Plugin) OnLoad() error {
	return nil
}

func (this *Plugin) OnReload() error {
	return nil
}

func (this *Plugin) OnStart() error {
	return nil
}

func (this *Plugin) OnStop() error {
	return nil
}

func (this *Plugin) OnUnload() error {
	return nil
}
