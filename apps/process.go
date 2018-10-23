package apps

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
	CPUUsage    *CPUUsage
	MemoryUsage *MemoryUsage

	OpenFiles   []string
	Connections []string
	Listens     []*Listen
}

func NewProcess(pid int32) *Process {
	return &Process{
		Pid: pid,
	}
}

// open files, connections, ...
func (this *Process) StatOpenFiles() {
	for _, result := range Lsof(this.Pid) {
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
