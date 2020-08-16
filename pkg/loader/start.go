package loader

import (
	"github.com/TeaWeb/plugin/pkg/plugins"
)

var sharedLoader *Loader

func Start(plugin *plugins.Plugin) {
	sharedLoader = NewLoader(plugin)
	sharedLoader.Load()
}
