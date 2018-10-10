package list

type DefaultFormatter struct {
}

func (f *DefaultFormatter) FormatTable(table Table, headerFormatter HeaderFormatter, rowFormatter RowFormatter) {
	panic("implement me")
}

func (f *DefaultFormatter) FormatRow(row Row, cellFormatter CellFormatter) (result string, err error) {
	panic("implement me")
}

func (f *DefaultFormatter) FormatHeader(cols []*Column, cellFormatter CellFormatter) (result string, err error) {
	for _, col := range cols {
	}
}

func (f *DefaultFormatter) FormatCell(value Value) (result string, err error) {
	return value, nil
}
