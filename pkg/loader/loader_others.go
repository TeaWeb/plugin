// +build !windows

package loader

import (
	"github.com/TeaWeb/plugin/pkg/messages"
	"github.com/TeaWeb/plugin/pkg/plugins"
	"os"
	"reflect"
)

func NewLoader(plugin *plugins.Plugin) *Loader {
	loader := &Loader{
		plugin:  plugin,
		methods: map[string]reflect.Method{},
	}

	loader.reader = os.NewFile(uintptr(3), "parentReader")
	loader.writer = os.NewFile(uintptr(4), "parentWriter")

	// 当前methods
	t := reflect.TypeOf(loader)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		loader.methods[method.Name] = method
	}

	loader.thisValue = reflect.ValueOf(loader)

	plugin.OnReloadedApps(func() {
		loader.Write(&messages.ReloadAppsAction{
			Apps: plugin.Apps,
		})
	})

	return loader
}
