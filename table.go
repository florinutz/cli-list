package list

// List aggregates all List components
type List struct {
	Columns []*Column
	Data    [][]string
}

func (l *List) WriteRow(incomingValues map[*Column]interface{}) (listDiff, inputDiff []*Column, err error) {
	inputDiff, listDiff = getColDiffs(incomingValues, l.Columns)

	return listDiff, inputDiff, nil
}

func getColDiffs(incomingValues map[*Column]interface{}, listCols []*Column) (inputDiff, listDiff []*Column) {
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
			if inputCol.Id == listCol.Id {
				valid = false
			}
		}
		if !valid {
			listDiff = append(listDiff, listCol)
		}
	}

	return
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

type RowWriter interface {
	WriteRow(values map[*Column]interface{}) (listDiff, inputDiff []*Column, err error)
}
