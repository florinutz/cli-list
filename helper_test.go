package list

import (
	"reflect"
	"testing"
)

func Test_getColDiffs(t *testing.T) {
	type args struct {
		incomingValues Row
		listCols       []*Column
	}
	tests := []struct {
		name          string
		args          args
		wantInputDiff []*Column
		wantListDiff  []*Column
	}{
		{
			"valid",
			args{
				Row{
					Data: map[*Column]interface{}{
						&Column{1, "id", nil}: interface{}("val"),
					},
					Formatter: nil,
				},
				[]*Column{
					{1, "id", nil},
				},
			},
			nil,
			nil,
		},
		{
			"diff",
			args{
				Row{
					Data: map[*Column]interface{}{
						&Column{1, "id", nil}: interface{}("val"),
					},
					Formatter: nil,
				},
				[]*Column{
					{2, "id", nil},
				},
			},
			[]*Column{{1, "id", nil}},
			[]*Column{{2, "id", nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInputDiff, gotListDiff := getColDiffs(tt.args.incomingValues, tt.args.listCols)
			if !reflect.DeepEqual(gotInputDiff, tt.wantInputDiff) {
				t.Errorf("getColDiffs() gotInputDiff = %v, want %v", gotInputDiff, tt.wantInputDiff)
			}
			if !reflect.DeepEqual(gotListDiff, tt.wantListDiff) {
				t.Errorf("getColDiffs() gotListDiff = %v, want %v", gotListDiff, tt.wantListDiff)
			}
		})
	}
}
