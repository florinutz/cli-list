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
					&Column{1, "id"}: interface{}("val"),
				},
				[]*Column{
					{1, "id"},
				},
			},
			nil,
			nil,
		},
		{
			"diff",
			args{
				Row{
					&Column{1, "id"}: interface{}("val"),
				},
				[]*Column{
					{2, "id"},
				},
			},
			[]*Column{{1, "id"}},
			[]*Column{{2, "id"}},
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
