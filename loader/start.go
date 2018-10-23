package loader

import (
	"github.com/TeaWeb/plugin/messages"
	"github.com/TeaWeb/plugin/plugins"
)

var sharedLoader *Loader

func Start(plugin *plugins.Plugin) {
	sharedLoader = NewLoader(plugin)
	sharedLoader.Load()
}

func ReloadApps(plugin *plugins.Plugin) {
	sharedLoader.Write(&messages.ReloadAppsAction{
		Apps: plugin.Apps,
	})
}
