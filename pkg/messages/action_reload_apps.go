package messages

import "github.com/TeaWeb/plugin/pkg/apps"

type ReloadAppsAction struct {
	Action

	Apps []*apps.App
}

func (this *ReloadAppsAction) Name() string {
	return "ReloadApps"
}
