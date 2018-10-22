package charts

type ProgressBar struct {
	Chart

	Value float64
	Color Color
}

func NewProgressBar() *ProgressBar {
	p := &ProgressBar{
		Color: ColorBlue,
	}
	p.ChartType = "progressBar"
	return p
}
