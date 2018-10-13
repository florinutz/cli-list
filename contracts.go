package list

import "io"

type TableFormatter interface {
	FormatTable(w io.Writer, table Table, headerFormatter HeaderFormatter, rowFormatter RowFormatter,
		cellFormatter CellFormatter) (headerErrs, rowsErrs []error, err error)
}

type HeaderFormatter interface {
	FormatHeader(w io.Writer, cols []*Column, cellFormatter CellFormatter) (errs []error)
}

type RowFormatter interface {
	FormatRow(w io.Writer, row Row, cellFormatter CellFormatter) (errs []error)
}

// Writer can write a row to something ( preferably a list ;) )
type RowWriter interface {
	Write(row Row) (listColsDiff, inputColsDiff []*Column, err error)
}

type CellFormatter interface {
	FormatCell(value interface{}) (result string, err error)
}
