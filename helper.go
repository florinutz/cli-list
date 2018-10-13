package list

import "strings"

func getColDiffs(incomingValues Row, listCols []*Column) (inputDiff, listDiff []*Column) {
	for inputCol := range incomingValues.Data {
		valid := false
		for _, listCol := range listCols {
			if inputCol.Id > 0 && inputCol.Id == listCol.Id {
				valid = true
			}
		}
		if !valid {
			inputDiff = append(inputDiff, inputCol)
		}
	}

	for _, listCol := range listCols {
		valid := false
		for inputCol := range incomingValues.Data {
			if inputCol.Id > 0 && inputCol.Id == listCol.Id {
				valid = true
			}
		}
		if !valid {
			listDiff = append(listDiff, listCol)
		}
	}

	return
}

func getTabwriterFormat(inputLen int) string {
	var s []string

	for i := 0; i < inputLen; i++ {
		s = append(s, "%s")
	}

	return strings.Join(s, "\t") + "\t\n"
}
