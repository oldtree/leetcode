package stringToIntegerAtoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"right",
			args{
				"+1",
			},
			1,
		},
		{
			"right",
			args{
				"-1",
			},
			-1,
		},
		{
			"right",
			args{
				"4193 with word",
			},
			4193,
		},
		{
			"right",
			args{
				"   -42",
			},
			-42,
		},
		{
			"right",
			args{
				"   42",
			},
			42,
		}, {
			"right",
			args{
				"words and 987",
			},
			0,
		}, {
			"right",
			args{
				"-91283472332",
			},
			-2147483648,
		}, {
			"right",
			args{
				"-1",
			},
			-1,
		},
		{
			"right",
			args{
				"-1",
			},
			-1,
		},
		{
			"right",
			args{
				"+-2",
			},
			0,
		},
		{
			"right",
			args{
				"-+2",
			},
			0,
		},
		{
			"right",
			args{
				"9223372036854775808",
			},
			2147483647,
		},
		{
			"right",
			args{
				"-5-",
			},
			-5,
		},
		{
			"right",
			args{
				"+5-",
			},
			5,
		},
		//"-9223372036854775809"
		{
			"right",
			args{
				"-9223372036854775809",
			},
			-2147483648,
		},
		{
			"right",
			args{
				"18446744073709551617",
			},
			2147483647,
		},
		{
			"right",
			args{
				"-18446744073709551617",
			},
			-2147483648,
		},
		{
			"right",
			args{
				"-2147483648",
			},
			-2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
