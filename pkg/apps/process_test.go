package apps

import (
	"testing"
)

func TestProcess(t *testing.T) {
	processes, _ := PsLookup("php-fpm", []string{}, true)
	for _, p := range processes {
		t.Log("====")
		t.Log(p.Pid, "parent:", p.Ppid, "cmd:", p.Cmdline)
		t.Log(p.CPUUsage)
		t.Log(p.MemoryUsage)
		t.Log("time:", p.CreateTime)

		p.StatOpenFiles()
		t.Log("connections:", p.Connections)
		for _, listen := range p.Listens {
			t.Log("listen:", listen)
		}
		t.Log("files:", p.OpenFiles)
	}
}
