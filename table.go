package list

import "github.com/pkg/errors"

// Table aggregates all Table components
type Table struct {
	cols []*Column
	rows []Row

	Writer    RowWriter
	Formatter TableFormatter
}

// Column represents a Table column
type Column struct {
	Id        int
	Label     string
	Formatter CellFormatter
}

// Row is a row
type Row struct {
	Data      map[*Column]interface{}
	Formatter RowFormatter
}

// NewTable instantiates a Table
func NewTable(cols []*Column, initialRows []Row, writer RowWriter, formatter TableFormatter) *Table {
	return &Table{
		Writer:    writer,
		Formatter: formatter,
		cols:      cols,
		rows:      initialRows,
	}
}

func (t *Table) Write(row Row) (listDiff, inputDiff []*Column, err error) {
	inputDiff, listDiff = getColDiffs(row, t.cols)
	if len(inputDiff) > 0 {
		err = errors.New("unknown input cols")
		return
	}

	t.rows = append(t.rows, row)

	return
}

// GetColumnNames returns column names
func (t *Table) GetColumnNames() (names []string) {
	for _, col := range t.cols {
		names = append(names, col.Label)
	}

	return
}

func (t *Table) String() string {
	//representation, err := t.Formatter.FormatTable(t)
}
