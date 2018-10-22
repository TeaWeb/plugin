package loader

import "github.com/TeaWeb/plugin/plugins"

func Start(plugin *plugins.Plugin) {
	NewLoader(plugin).Load()
}
