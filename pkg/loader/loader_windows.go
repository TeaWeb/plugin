// +build windows

package loader

import (
	"github.com/Microsoft/go-winio"
	"github.com/TeaWeb/plugin/pkg/messages"
	"github.com/TeaWeb/plugin/pkg/plugins"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

func NewLoader(plugin *plugins.Plugin) *Loader {
	filename := filepath.Base(os.Args[0])
	rFile := `\\.\pipe\teaweb.reader.` + filename + `.pipe`
	wFile := `\\.\pipe\teaweb.writer.` + filename + `.pipe`

	loader := &Loader{
		plugin:  plugin,
		methods: map[string]reflect.Method{},
	}

	rConn, err := winio.DialPipe(wFile, nil)
	if err != nil {
		log.Println("[plugin]dial reader pipe:" + err.Error())
	} else {
		loader.reader = rConn
	}

	wConn, err := winio.DialPipe(rFile, nil)
	if err != nil {
		log.Println("[plugin]dial writer pipe:" + err.Error())
	} else {
		loader.writer = wConn
	}

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
