package messages

import "github.com/TeaWeb/plugin/pkg/apps"

type ReloadAppAction struct {
	Action

	App *apps.App
}

func (this *ReloadAppAction) Name() string {
	return "ReloadApp"
}
