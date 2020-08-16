package probes

import (
	"github.com/TeaWeb/plugin/pkg/apps"
	"strings"
)

type ProcessProbe struct {
	Probe

	CommandName     string   // 命令名
	CommandPatterns []string // 命令行匹配模式

	// 获取版本号信息的命令
	// 支持以下变量：
	// - ${commandFile} 命令行文件路径
	// - ${commandDir} 命令行文件所在目录
	CommandVersion string

	versionParser     func(versionString string) (string, error) // 版本文本分析器，同CommandVersion配合使用
	processValidators []func(proc *apps.Process) bool            // 进程校验
}

func NewProcessProbe() *ProcessProbe {
	return &ProcessProbe{}
}

func (this *ProcessProbe) Run() (resultApps []*apps.App, err error) {
	processes, err := apps.PsLookup(this.CommandName, this.CommandPatterns, true)
	if err != nil {
		return
	}

MainLoop:
	for _, mainProcess := range processes {
		// 校验进程
		for _, v := range this.processValidators {
			if !v(mainProcess) {
				continue MainLoop
			}
		}

		// 构建新App
		app := new(apps.App)
		app.Name = this.Name
		app.Site = this.Site
		app.DocSite = this.DocSite
		app.Developer = this.Developer
		app.IsRunning = true
		app.AddAllProcess(mainProcess)
		app.Cmdline = mainProcess.Cmdline

		// 版本号
		if len(this.CommandVersion) > 0 {
			// args
			versionCmd := strings.Replace(this.CommandVersion, "${commandDir}", "\""+mainProcess.Dir+"\"", -1)
			versionCmd = strings.Replace(versionCmd, "${commandFile}", "\""+mainProcess.File+"\"", -1)

			args := apps.ParseArgs(versionCmd)
			output, err := apps.Exec(args[0], args[1:] ...)
			if err == nil {
				if this.versionParser != nil {
					output, err := this.versionParser(output)
					if err == nil {
						app.Version = output
					}
				} else {
					app.Version = output
				}
			}
		}

		// 重新尝试检测
		app.OnStop(func() error {
			processes, err := apps.PsLookup(this.CommandName, this.CommandPatterns, true)
			if err != nil {
				return nil
			}

			for _, p := range processes {
				if p.Cmdline == app.Cmdline {
					app.ResetProcesses()
					app.AddAllProcess(p)
				}
			}

			return nil
		})

		resultApps = append(resultApps, app)
	}

	return
}

func (this *ProcessProbe) OnProcess(validator func(process *apps.Process) bool) {
	this.processValidators = append(this.processValidators, validator)
}

func (this *ProcessProbe) OnParseVersion(parser func(versionString string) (string, error)) {
	this.versionParser = parser
}
