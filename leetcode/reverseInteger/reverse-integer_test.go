package reverseInteger

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"right",
			args{
				123,
			},
			321,
		},
		{
			"right",
			args{
				222,
			},
			222,
		},
		{
			"right",
			args{
				10,
			},
			1,
		},
		{
			"right",
			args{
				101,
			},
			101,
		},
		{
			"right",
			args{
				100000,
			},
			1,
		},
		{
			"right",
			args{
				-100000,
			},
			-1,
		},
		{
			"right",
			args{
				-12345,
			},
			-54321,
		},
		{
			"right",
			args{
				-101,
			},
			-101,
		},
		{
			"right",
			args{
				0,
			},
			0,
		},
		{
			"right",
			args{
				-1,
			},
			-1,
		},
		{
			"right",
			args{
				120,
			},
			21,
		},
		{
			"right",
			args{
				-120,
			},
			-21,
		},
		{
			"32",
			args{
				1534236469,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
