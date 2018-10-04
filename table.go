package list

// List aggregates all List components
type List struct {
	Columns []*Column
	Data    [][]string
}

// Column represents a List column
type Column struct {
	Name string
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
		names = append(names, col.Name)
	}

	return
}
