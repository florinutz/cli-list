package test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/florinutz/cli-list"
)

func TestList_EndToEnd(t *testing.T) {
	type fields struct {
		Columns      []*list.Column
		DataProvider list.DataProvider
	}

	type args struct {
		viewer list.Flattener
		hi     list.StringHighlighterFunc
		quiet  bool
	}

	listInput := fields{
		[]*list.Column{
			{Label: "first"},
			{Label: "second"},
		},
		list.MemoryDataProviderFunc(func() [][]string {
			return [][]string{
				{"foo1", "bar1"},
				{"foo2", "bar2"},
			}
		}),
	}

	highlighterFunc := list.StringHighlighterFunc(color.New(color.FgGreen).SprintFunc())

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "default flattener",
			fields: listInput,
			args: args{
				viewer: list.DefaultListFlattener,
				hi:     highlighterFunc,
				quiet:  false,
			},
			want: `first        second
foo1         bar1
foo2         bar2
`,
			wantErr: false,
		},
		{
			name:   "csv flattener",
			fields: listInput,
			args: args{
				viewer: list.CsvListFlattener,
				hi:     highlighterFunc,
				quiet:  false,
			},
			want: `foo1,bar1
foo2,bar2
`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.NewList(tt.fields.Columns, tt.fields.DataProvider)
			got, err := l.Flatten(tt.args.viewer, tt.args.hi, tt.args.quiet)
			if (err != nil) != tt.wantErr {
				t.Errorf("List.Flatten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("List.Flatten() = \n%v\n want \n\n%v", got, tt.want)
			}
		})
	}
}
