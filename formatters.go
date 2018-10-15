package list

import (
	"encoding/csv"
	"io"
	"reflect"

	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/pkg/errors"
)

type Formatter interface {
	Format(writer io.Writer, data [][]string) error
}

type FormatterFunc func(writer io.Writer, data [][]string) error

func (f FormatterFunc) Format(writer io.Writer, data [][]string) error {
	return f(writer, data)
}

// CsvFormatter flattens to CSV
var CsvFormatter FormatterFunc = func(writer io.Writer, data [][]string) error {
	if len(data) == 0 {
		return nil
	}
	w := csv.NewWriter(writer)
	err := w.WriteAll(data)
	if err != nil {
		return errors.Wrap(err, "couldn't write data to csv buffer")
	}

	return nil
}

var TabsFormatter FormatterFunc = func(writer io.Writer, data [][]string) error {
	var DLFTabWriter = tabwriter.NewWriter(writer, 0, 0, 8, ' ', 0)

	// write data
	for _, row := range data {
		// todo GetData should store interface{} values instead of strings
		format := getTabwriterFormat(len(row))

		is, err := interfaceSlice(row)
		if err != nil {
			return errors.Wrap(err, "error converting strings to interfaces")
		}

		fmt.Fprintf(DLFTabWriter, format, is...)
	}

	err := DLFTabWriter.Flush()
	if err != nil {
		return errors.Wrap(err, "couldn't flush the tab writer")
	}

	return nil
}

type TableFormatter struct {
	Columns []string
}

func (f *TableFormatter) Format(writer io.Writer, data [][]string) error {
	if len(data) == 0 || len(data[0]) != len(f.Columns) {
		return fmt.Errorf("the number of data columns (%d) doest't correspond "+
			"to the number of table cols (%d)", len(data[0]), len(f.Columns))
	}
	newData := append([][]string{f.Columns}, data...)

	return TabsFormatter(writer, newData)
}

func getTabwriterFormat(inputLen int) string {
	var s []string

	for i := 0; i < inputLen; i++ {
		s = append(s, "%s")
	}

	return strings.Join(s, "\t") + "\t\n"
}

func interfaceSlice(slice interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, nil
}
