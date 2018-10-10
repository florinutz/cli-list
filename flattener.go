package list

// StringHighlighterFunc is the type of a function that highlights a string
type StringHighlighterFunc func(a ...interface{}) string

// FlattenerFunc is a function that converts a Table to string.
// If Quiet is specified, a short version should be returned, typically the first column only.
type FlattenerFunc func(list Table, hi StringHighlighterFunc, quiet bool) (string, error)

// FlattenList implements the flattener interface
func (f FlattenerFunc) FlattenList(list Table, hi StringHighlighterFunc, quiet bool) (string, error) {
	return f(list, hi, quiet)
}
