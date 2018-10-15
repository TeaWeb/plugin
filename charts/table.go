package charts

import "sync"

type Row struct {
	columns []interface{}
}

func (this *Row) Columns() []interface{} {
	return this.columns
}

func (this *Row) SetColumns(columns []interface{}) {
	this.columns = columns
}

type Column struct {
	text  string
	width float64 // 百分比，比如 30 表示 30%
}

func (this *Column) Text() string {
	return this.text
}

func (this *Column) SetText(text string) {
	this.text = text
}

func (this *Column) SetWidth(width float64) {
	this.width = width
}

func (this *Column) Width() float64 {
	return this.width
}

func NewTable() *Table {
	p := &Table{
		rows: []interface{}{},
	}
	p.SetType("table")
	return p
}

type Table struct {
	Chart

	rows   []interface{}
	locker sync.Mutex

	width []float64
}

func (this *Table) Rows() []interface{} {
	return this.rows
}

func (this *Table) ResetRows() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.rows = []interface{}{}
}

func (this *Table) AddRow(text ... string) {
	this.locker.Lock()
	defer this.locker.Unlock()

	columns := []interface{}{}
	for index, t := range text {
		if index < len(this.width) {
			columns = append(columns, &Column{
				text:  t,
				width: this.width[index],
			})
		} else {
			columns = append(columns, &Column{
				text: t,
			})
		}
	}
	this.rows = append(this.rows, &Row{
		columns: columns,
	})
}

func (this *Table) SetWidth(wide ... float64) {
	this.width = wide

	for _, row := range this.rows {
		for index, column := range row.(*Row).columns {
			if index < len(this.width) {
				column.(*Column).width = this.width[index]
			}
		}
	}
}
