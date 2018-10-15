package charts

import "github.com/TeaWeb/plugin/teainterfaces"

type ProgressBar struct {
	Chart

	value float64
	color teainterfaces.Color
}

func NewProgressBar() *ProgressBar {
	p := &ProgressBar{
		color: teainterfaces.ColorBlue,
	}
	p.SetType("progressBar")
	return p
}

func (this *ProgressBar) SetValue(value float64) {
	this.value = value
}

func (this *ProgressBar) Value() float64 {
	return this.value
}

func (this *ProgressBar) SetColor(color teainterfaces.Color) {
	this.color = color
}

func (this *ProgressBar) Color() teainterfaces.Color {
	return this.color
}
