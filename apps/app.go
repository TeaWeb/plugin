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
	Operations []*Operation
	Monitors   []*Monitor
	Statistics []*Statistics
	Logs       []*Log
}
