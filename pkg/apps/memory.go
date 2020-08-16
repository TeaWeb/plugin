package apps

// 内存使用
type MemoryUsage struct {
	RSS     uint64  // RSS
	VMS     uint64  // VMS
	Percent float64 // 百分比
}
