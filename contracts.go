package list

// Writer can write a row to something ( preferably a list ;) )
type RowWriter interface {
	Write(row Row) (listColsDiff, inputColsDiff []*Column, err error)
}

type TableFormatter interface {
	FormatTable(table Table, headerFormatter HeaderFormatter, rowFormatter RowFormatter) (
		result string, headerErrs, rowsErrs []error)
}

type HeaderFormatter interface {
	FormatHeader(cols []*Column, cellFormatter CellFormatter) (result string, errs []error)
}

type RowFormatter interface {
	FormatRow(row Row, cellFormatter CellFormatter) (result string, errs []error)
}

type CellFormatter interface {
	FormatCell(value interface{}) (result string, err error)
}
