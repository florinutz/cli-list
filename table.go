package list

import "github.com/pkg/errors"

// Table aggregates all Table components
type Table struct {
	Writer    RowWriter
	Formatter TableFormatter

	columns []*Column
	rows    []Row
}

// Column represents a Table column
type Column struct {
	Id    int
	Label string
}

// Row is a row
type Row map[*Column]Value

// Writer can write a row to something ( preferably a list ;) )
type RowWriter interface {
	Write(row Row) (listColsDiff, inputColsDiff []*Column, err error)
}

type HeaderFormatter interface {
	FormatHeader(cols []*Column, cellFormatter CellFormatter) (result string, err error)
}

type RowFormatter interface {
	FormatRow(row Row, cellFormatter CellFormatter) (result string, err error)
}

type CellFormatter interface {
	FormatCell(value string) (result string, err error)
}

type TableFormatter interface {
	FormatTable(table Table, headerFormatter HeaderFormatter, rowFormatter RowFormatter)
}

// Value represents a cell's value
type Value interface{}

// NewTable instantiates a Table
func NewTable(cols []*Column, writer RowWriter, formatter TableFormatter) *Table {
	return &Table{
		Writer:    writer,
		Formatter: formatter,
		columns:   cols,
	}
}

func (t *Table) Write(row Row) (listDiff, inputDiff []*Column, err error) {
	inputDiff, listDiff = getColDiffs(row, t.columns)
	if len(inputDiff) > 0 {
		err = errors.New("unknown input cols")
		return
	}

	t.rows = append(t.rows, row)

	return
}

// GetColumnNames returns column names
func (t *Table) GetColumnNames() (names []string) {
	for _, col := range t.columns {
		names = append(names, col.Label)
	}

	return
}

func (t *Table) String() string {
	//representation, err := t.Formatter.FormatTable(t)
}
