package messages

import "github.com/TeaWeb/plugin/apps"

type ReloadAppAction struct {
	Action

	App *apps.App
}

func (this *ReloadAppAction) Name() string {
	return "ReloadApp"
}
