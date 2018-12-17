package plusOne

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digit []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"right",
			args{
				[]int{1, 2, 3},
			},
			[]int{1, 2, 4},
		},
		{
			"right",
			args{
				[]int{1, 2, 9},
			},
			[]int{1, 3, 0},
		},
		{
			"right",
			args{
				[]int{9, 9, 9},
			},
			[]int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
