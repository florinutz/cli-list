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
	FormatHeader(cols []*Column) (result string, err error)
}

type RowsFormatter interface {
	FormatRow(rows []Row, cellFormatter CellFormatter) (result string, err error)
}

type CellFormatter interface {
	FormatCell(value Value) (result string, err error)
}

type TableFormatter interface {
	FormatTable(table Table, headerFormatter HeaderFormatter, rowsFormatter RowsFormatter)
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

func (t *Table) Write(incomingRow Row) (listDiff, inputDiff []*Column, err error) {
	inputDiff, listDiff = getColDiffs(incomingRow, l.columns)
	if len(inputDiff) > 0 {
		err = errors.New("unknown input cols")
		return
	}

	l.rows = append(l.rows, incomingRow)

	return
}

func getColDiffs(incomingValues Row, listCols []*Column) (inputDiff, listDiff []*Column) {
	for inputCol := range incomingValues {
		valid := false
		for _, listCol := range listCols {
			if inputCol.Id > 0 && inputCol.Id == listCol.Id {
				valid = false
			}
		}
		if !valid {
			inputDiff = append(inputDiff, inputCol)
		}
	}

	for _, listCol := range listCols {
		valid := false
		for inputCol := range incomingValues {
			if inputCol.Id > 0 && inputCol.Id == listCol.Id {
				valid = false
			}
		}
		if !valid {
			listDiff = append(listDiff, listCol)
		}
	}

	return
}

// GetColumnNames returns column names
func (t *Table) GetColumnNames() (names []string) {
	for _, col := range l.columns {
		names = append(names, col.Label)
	}

	return
}

func (t *Table) String() string {
	panic("implement me")
}
