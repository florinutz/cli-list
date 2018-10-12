package list

// DefaultListFlattener provides a default flattener with tabs
var DefaultListFlattener FlattenerFunc = func(l Table, hi StringHighlighterFunc, quiet bool) (string, error) {
	return "", nil
	// var b strings.Builder
	// var DLFTabWriter = tabwriter.NewWriter(&b, 0, 0, 8, ' ', 0)
	//
	// if len(l.GetData()) == 0 {
	// 	return "", nil
	// }
	//
	// if !quiet && len(l.cols) > 0 {
	// 	format := getTabwriterFormat(len(l.cols))
	//
	// 	is, err := interfaceSlice(l.GetColumnNames())
	// 	if err != nil {
	// 		return "", errors.Wrap(err, "error converting strings to interfaces")
	// 	}
	//
	// 	fmt.Fprintf(DLFTabWriter, format, is...)
	// }
	//
	// // write data
	// for _, row := range l.GetData() {
	// 	if quiet && len(row) > 0 {
	// 		row = row[:1]
	// 	}
	//
	// 	// todo GetData should store interface{} values instead of strings
	// 	format := getTabwriterFormat(len(row))
	//
	// 	is, err := interfaceSlice(row)
	// 	if err != nil {
	// 		return "", errors.Wrap(err, "error converting strings to interfaces")
	// 	}
	//
	// 	fmt.Fprintf(DLFTabWriter, format, is...)
	// }
	//
	// err := DLFTabWriter.Flush()
	// if err != nil {
	// 	return "", errors.Wrap(err, "couldn't flush the tab writer")
	// }
	//
	// return b.String(), nil
}
