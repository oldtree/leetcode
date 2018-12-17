package monotonicArray

import "testing"

func Test_monotonicArray(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"right",
			args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			true,
		},
		{
			"right",
			args{
				array: []int{1, 2, 3},
			},
			true,
		},
		{
			"wrong",
			args{
				array: []int{1, 3, 2},
			},
			false,
		},
		{
			"wrong",
			args{
				array: []int{3, 3, 3},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMonotonic(tt.args.array); got != tt.want {
				t.Errorf("monotonicArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
