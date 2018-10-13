package list

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/pkg/errors"
)

type TabwriterFormatter struct {
	*tabwriter.Writer
	builder *strings.Builder
}

func NewTabwriterFormatter(
	minwidth, tabwidth, padding int, padchar byte, flags uint) *TabwriterFormatter {
	builder := &strings.Builder{}
	return &TabwriterFormatter{
		Writer:  tabwriter.NewWriter(builder, minwidth, tabwidth, padding, padchar, flags),
		builder: builder,
	}
}

func (df *TabwriterFormatter) Write(w io.Writer, values []interface{}) error {
	if len(values) == 0 {
		return errors.New("empty input")
	}

	format := getTabwriterFormat(len(values))
	if _, err := fmt.Fprintf(w, format, values...); err != nil {
		return err
	}

	return nil
}

func (df *TabwriterFormatter) FormatTable(
	w io.Writer, table Table, headerFormatter HeaderFormatter, rowFormatter RowFormatter, cellFormatter CellFormatter) (
	headerErrs, rowsErrs []error, err error) {
	if headerFormatter == nil {
		headerFormatter = df
	}
	if rowFormatter == nil {
		rowFormatter = df
	}

	headerErrs = headerFormatter.FormatHeader(df.Writer, table.cols, cellFormatter)

	for _, row := range table.rows {
		rowErrs := rowFormatter.FormatRow(df.Writer, row, cellFormatter)
		if len(rowErrs) > 0 {
			rowsErrs = append(rowsErrs, rowErrs...)
		}
	}

	if err = df.Writer.Flush(); err != nil {
		return
	}

	return
}

func (df *TabwriterFormatter) FormatRow(
	w io.Writer, row Row, cellFormatter CellFormatter) (errs []error) {
	if cellFormatter == nil {
		cellFormatter = df
	}
	var values []interface{}
	for col, val := range row.Data {
		str, err := cellFormatter.FormatCell(val)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "error formatting cell of column '%s'", col.Label))
		} else {
			values = append(values, str)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	if err := df.Write(w, values); err != nil {
		return []error{err}
	}

	return nil
}

func (df *TabwriterFormatter) FormatHeader(w io.Writer, cols []*Column, cellFormatter CellFormatter) (errs []error) {
	var labels []interface{}

	for _, col := range cols {
		if cellFormatter == nil {
			cellFormatter = col.Formatter
		}

		var str string
		if cellFormatter != nil {
			aux, err := cellFormatter.FormatCell(col.Label)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			str = aux
		} else {
			str = col.Label
		}
		labels = append(labels, str)
	}
	if len(errs) > 0 {
		return errs
	}

	if err := df.Write(w, labels); err != nil {
		return []error{err}
	} else {
		return nil
	}
}

func (df *TabwriterFormatter) FormatCell(value interface{}) (result string, err error) {
	var valString string

	if valStr, ok := value.(string); !ok {
		if valStringer, ok2 := value.(fmt.Stringer); !ok2 {
			return "", errors.New("can't convert value to string")
		} else {
			valString = valStringer.String()
		}
	} else {
		valString = valStr
	}

	return valString, nil
}
