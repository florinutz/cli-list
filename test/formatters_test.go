package test

import (
	"bytes"
	"testing"

	"github.com/florinutz/cli-list"
)

func TestFormatterFunc_Format(t *testing.T) {
	type args struct {
		data [][]string
	}
	data := [][]string{
		{"one", "two", "three"},
		{"th ree", "fou,r", "five"},
	}

	tests := []struct {
		name       string
		f          list.Formatter
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "csv",
			f:    list.CsvFormatter,
			args: args{data: data},
			wantOutput: `one,two,three
th ree,"fou,r",five
`,
			wantErr: false,
		},
		{
			name: "tabs",
			f:    list.TabsFormatter,
			args: args{data: data},
			wantOutput: `one           two          three        
th ree        fou,r        five         
`,
			wantErr: false,
		},
		{
			name: "table",
			f:    &list.TableFormatter{Columns: []string{"col1", "col2", "col3"}},
			args: args{data: data},
			wantOutput: `col1          col2         col3         
one           two          three        
th ree        fou,r        five         
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := tt.f.Format(writer, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("FormatterFunc.Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantOutput {
				t.Errorf("FormatterFunc.Format() = \n%v\n\n, want \n%v\n", gotWriter, tt.wantOutput)
			}
		})
	}
}
