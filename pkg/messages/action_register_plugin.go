package messages

import "github.com/TeaWeb/plugin/pkg/plugins"

type RegisterPluginAction struct {
	Action

	Plugin *plugins.Plugin
}

func (this *RegisterPluginAction) Name() string {
	return "RegisterPlugin"
}
