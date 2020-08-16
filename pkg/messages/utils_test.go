package messages

import (
	"github.com/TeaWeb/plugin/pkg/plugins"
	"testing"
)

func TestNewAction(t *testing.T) {
	p1 := NewAction("RegisterPlugin").(*RegisterPluginAction)
	p1.Plugin = &plugins.Plugin{}

	p2 := NewAction("RegisterPlugin").(*RegisterPluginAction)
	p2.Plugin = &plugins.Plugin{}
	t.Logf("%p, %p", p1, p2)
}
