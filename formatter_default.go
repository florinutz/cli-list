package list

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/pkg/errors"
)

type defaultFormatter struct {
	tabwriterOptions tabwriterOptions
}

type tabwriterOptions struct {
	minwidth int
	tabwidth int
	padding  int
	padchar  byte
	flags    uint
}

func NewDefaultFormatter(minwidth, tabwidth, padding int, padchar byte, flags uint) *defaultFormatter {
	return &defaultFormatter{
		tabwriterOptions: tabwriterOptions{
			minwidth: minwidth,
			tabwidth: tabwidth,
			padding:  padding,
			padchar:  padchar,
			flags:    flags,
		},
	}
}

func (df *defaultFormatter) getString(values []interface{}) (string, error) {
	if len(values) == 0 {
		return "", errors.New("empty input")
	}

	var b strings.Builder
	tw := tabwriter.NewWriter(&b, 0, 0, 8, ' ', 0)

	format := df.getTabwriterFormat(len(values))
	if _, err := fmt.Fprintf(tw, format, values...); err != nil {
		return "", err
	}

	if err := tw.Flush(); err != nil {
		return "", errors.New("tabwriter flush error")
	}

	return b.String(), nil
}

func (df *defaultFormatter) FormatTable(table Table, headerFormatter HeaderFormatter, rowFormatter RowFormatter) (
	result string, headerErrs, rowsErrs []error) {
	headerString, headerErrs := headerFormatter.FormatHeader(table.cols, nil)
	if len(headerErrs) > 0 {
	}

	var rowStrings []string
	for _, row := range table.rows {
		rowStr, rowErrs := rowFormatter.FormatRow(row, nil)
		if len(rowErrs) > 0 {
			rowsErrs = append(rowsErrs, rowErrs...)
		} else {
			rowStrings = append(rowStrings, rowStr)
		}
	}
	rowsString := strings.Join(rowStrings, "\n")

	result = fmt.Sprintf("%s\n%s", headerString, rowsString)

	return
}

func (df *defaultFormatter) FormatRow(row Row, cellFormatter CellFormatter) (result string, errs []error) {
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
		return "", errs
	}

	str, err := df.getString(values)
	if err != nil {
		return "", []error{err}
	}

	return str, nil
}

func (df *defaultFormatter) FormatHeader(cols []*Column, cellFormatter CellFormatter) (result string, errs []error) {
	var labels []interface{}

	for _, col := range cols {
		var f CellFormatter

		if col.Formatter != nil {
			f = col.Formatter
		} else {
			f = cellFormatter
		}

		str, err := f.FormatCell(col.Label)
		if err != nil {
			errs = append(errs, err)
		} else {
			labels = append(labels, str)
		}
	}
	if len(errs) > 0 {
		return "", errs
	}

	if str, err := df.getString(labels); err != nil {
		return "", []error{err}
	} else {
		return str, nil
	}
}

func (df *defaultFormatter) FormatCell(value interface{}) (result string, err error) {
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

func (df *defaultFormatter) getTabwriterFormat(inputLen int) string {
	var s []string

	for i := 0; i < inputLen; i++ {
		s = append(s, "%s")
	}

	return strings.Join(s, "\t") + "\n"
}
