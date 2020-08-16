package apps

import (
	"strings"
	"testing"
)

func TestPs(t *testing.T) {
	// ps -p 58682 -o user,pid,ppid,%cpu,cputime,uid,gid,user,lstart,%mem,command

	processes, _ := PsLookup("mongodb", []string{"mongod$", "fork"}, false)
	if len(processes) == 0 {
		t.Log("mongo process not found")
		return
	}
	for _, proc := range processes {
		t.Log("pid:", proc.Pid)
		t.Log("name:", proc.Name)
		t.Log("cmdLine:", proc.Cmdline)
		t.Log("cwd:", proc.Cwd)
		t.Log(proc.CPUUsage)
		t.Log(proc.MemoryUsage)

		// username & group
		t.Log(proc.User)
		t.Log("uid:", proc.Uid)
		t.Log("gid:", proc.Gid)

		proc.StatOpenFiles()
		t.Log(proc.OpenFiles)
		t.Log(proc.Listens)
		t.Log(proc.Connections)
	}
}

func TestCmdSlice(t *testing.T) {
	cmd := "bin/mongod --dbpath=./data --fork --logpath=\"./data/fork.log\""
	t.Log(strings.Join(ParseArgs(cmd), "\n"))
}

func TestPsChildren(t *testing.T) {
	result, err := PsLookup("php-fpm", []string{}, true)
	if err != nil {
		t.Fatal(err)
	}
	pid := result[0].Pid
	children, err := PsChildren(pid)
	if err != nil {
		t.Fatal(err)
	}

	for _, child := range children {
		t.Log(child)
	}
}

func TestPsDir(t *testing.T) {
	result, err := PsLookup("httpd", []string{}, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) > 0 {
		t.Log(result[0])
	}
}
