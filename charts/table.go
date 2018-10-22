package charts

import "sync"

type Row struct {
	Columns []*Column
}

type Column struct {
	Text  string
	Width float64 // 百分比，比如 30 表示 30%
}

func NewTable() *Table {
	p := &Table{
		Rows: []*Row{},
	}
	p.ChartType = "table"
	return p
}

type Table struct {
	Chart

	Rows   []*Row
	locker sync.Mutex

	Width []float64
}

func (this *Table) ResetRows() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.Rows = []*Row{}
}

func (this *Table) AddRow(text ...string) {
	this.locker.Lock()
	defer this.locker.Unlock()

	columns := []*Column{}
	for index, t := range text {
		if index < len(this.Width) {
			columns = append(columns, &Column{
				Text:  t,
				Width: this.Width[index],
			})
		} else {
			columns = append(columns, &Column{
				Text: t,
			})
		}
	}
	this.Rows = append(this.Rows, &Row{
		Columns: columns,
	})
}

func (this *Table) SetWidth(wide ...float64) {
	this.Width = wide

	for _, row := range this.Rows {
		for index, column := range row.Columns {
			if index < len(this.Width) {
				column.Width = this.Width[index]
			}
		}
	}
}
