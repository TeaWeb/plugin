package apps

import (
	"path/filepath"
	"strings"
)

type Listen struct {
	Network string
	Addr    string
}

type Process struct {
	Name string
	Pid  int32
	Ppid int32
	Cwd  string

	User string
	Uid  int32
	Gid  int32

	CreateTime  int64 // 时间戳
	Cmdline     string
	File        string // 命令行文件路径
	Dir         string // 命令行文件所在目录
	CPUUsage    *CPUUsage
	MemoryUsage *MemoryUsage

	OpenFiles   []string
	Connections []string
	Listens     []*Listen

	IsRunning bool
}

func NewProcess(pid int32) *Process {
	return &Process{
		Pid:       pid,
		IsRunning: true,
	}
}

// 修改命令的名称
func (this *Process) ChangeName(newName string) {
	this.Name = newName

	args := ParseArgs(this.Cmdline)
	for _, arg := range args {
		if strings.HasSuffix(arg, "/"+this.Name) {
			this.File = arg
			if arg[0] == '/' {
				this.Dir = filepath.Dir(this.File)
			} else {
				this.File = this.Cwd + "/" + this.File
				absFile, err := filepath.Abs(this.File)
				if err == nil {
					this.File = absFile
				}
				this.Dir = filepath.Dir(this.File)
			}
			break
		}
	}
}

// open files, connections, ...
func (this *Process) StatOpenFiles() {
	results, err := Lsof(this.Pid)
	if err != nil {
		return
	}
	for _, result := range results {
		if result.IsCwd() {
			if len(this.Cwd) == 0 {
				this.Cwd = result.Name
			}
		} else if result.IsRegularFile() {
			this.OpenFiles = append(this.OpenFiles, result.Name)
		} else if result.IsListening() {
			this.Listens = append(this.Listens, &Listen{
				Network: result.Protocol,
				Addr:    result.Listen(),
			})
		} else if result.IsEstablished() {
			this.Connections = append(this.Connections, result.LAddr()+"->"+result.RAddr())
		}
	}
}

// 刷新状态
func (this *Process) Reload() error {
	this.CPUUsage = nil
	this.MemoryUsage = nil

	this.OpenFiles = []string{}
	this.Connections = []string{}
	this.Listens = []*Listen{}

	p, err := PsPid(this.Pid)
	if err != nil {
		this.IsRunning = false
		return err
	}

	if p.Cmdline != this.Cmdline || p.Ppid != this.Ppid {
		this.IsRunning = false
		return nil
	}

	this.IsRunning = true
	this.CPUUsage = p.CPUUsage
	this.MemoryUsage = p.MemoryUsage

	return nil
}
