package stats

import (
	"reflect"
	"testing"
)

func TestSortSlice(t *testing.T) {
	type args struct {
		transactions []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{name: "slice", args: args{[]int64{120, 234, 12, 122, 234, 794}}, want: []int64{794, 234, 234, 122, 120, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortSlice(tt.args.transactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}