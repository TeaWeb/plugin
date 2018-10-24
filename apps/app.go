package apps

// 接口
type AppInterface interface {
	Start() error
	Stop() error
}

// App定义
type App struct {
	Id        string // 唯一ID，通常系统会自动生成
	Name      string
	Developer string
	Site      string
	DocSite   string
	Version   string
	Icon      []byte

	Processes  []*Process
	Operations []*Operation  // TODO
	Monitors   []*Monitor    // TODO
	Statistics []*Statistics // TODO
	Logs       []*Log        // TODO

	IsRunning bool
}

// 重置进程
func (this *App) ResetProcesses() {
	this.Processes = []*Process{}
}

// 添加进程，不包括子进程
func (this *App) AddProcess(process ... *Process) {
	this.Processes = append(this.Processes, process ...)
}

// 添加进程，包括子进程
func (this *App) AddAllProcess(process ... *Process) {
	for _, p := range process {
		p.StatOpenFiles()
		this.AddProcess(p)

		children, err := PsChildren(p.Pid)
		if err != nil {
			continue
		}

		for _, c := range children {
			c.StatOpenFiles()
			this.AddProcess(c)
		}
	}
}
