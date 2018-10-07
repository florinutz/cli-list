package list

// List aggregates all List components
type List struct {
	Columns []*Column
	Data    [][]string
}

func (l *List) WriteRow(inputValues map[*Column]interface{}) error {
	var inputDiff, listDiff []*Column

	listCols := l.Columns

	for inputCol := range inputValues {
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
		for inputCol := range inputValues {
			if inputCol.Label == listCol.Label {
				valid = false
			}
		}
		if !valid {
			listDiff = append(listDiff, listCol)
		}
	}
}

// Column represents a List column
type Column struct {
	Id    int
	Label string
}

// DataProvider provides the data for a List
type DataProvider interface {
	GetData() [][]string
}

type RowWriter interface {
	WriteRow(values map[*Column]interface{}) error
}

// Flattener can display a List
type Flattener interface {
	FlattenList(list List, hi StringHighlighterFunc, quiet bool) (string, error)
}

// NewList instantiates a List
func NewList(cols []*Column, provider DataProvider) *List {
	return &List{
		Columns: cols,
		Data:    provider.GetData(),
	}
}

// Flatten flattens the List receiver to string
func (l *List) Flatten(viewer Flattener, hi StringHighlighterFunc, quiet bool) (string, error) {
	return viewer.FlattenList(*l, hi, quiet)
}

// GetData implements DataProvider
func (l *List) GetData() [][]string {
	return l.Data
}

// GetColumnNames returns column names
func (l *List) GetColumnNames() (names []string) {
	for _, col := range l.Columns {
		names = append(names, col.Label)
	}

	return
}
