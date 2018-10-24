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

	Cmdline   string

	Processes  []*Process    // 进程列表
	Operations []*Operation  // TODO
	Monitors   []*Monitor    // TODO
	Statistics []*Statistics // TODO
	Logs       []*Log        // TODO

	IsRunning bool

	onStopFunc   func() error
	onReloadFunc func() error
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

// 设置reload回调
func (this *App) OnReload(f func() error) {
	this.onReloadFunc = f
}

// 刷新App
func (this *App) Reload() error {
	// 刷新子进程
	if len(this.Processes) > 0 {
		mainProcess := this.Processes[0]
		mainProcess.Reload()

		this.IsRunning = mainProcess.IsRunning

		this.ResetProcesses()
		if mainProcess.IsRunning {
			this.AddAllProcess(mainProcess)
		}
	}

	if !this.IsRunning {
		if this.onStopFunc != nil {
			this.onStopFunc()
		}
	}

	if this.onReloadFunc != nil {
		return this.onReloadFunc()
	}
	return nil
}

// 设置进程停止时回调
func (this *App) OnStop(f func() error) {
	this.onStopFunc = f
}
