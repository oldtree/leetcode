package longestPalindrome

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"right",
			args{
				"sssss",
			},
			"ss",
		},
		{
			"right",
			args{
				"babad",
			},
			"bab",
		},
		{
			"right",
			args{
				"cbbd",
			},
			"bb",
		},
		{
			"right",
			args{
				"bbbbbb",
			},
			"bb",
		},
		{
			"right",
			args{
				"b",
			},
			"b",
		},
		{
			"right",
			args{
				"abc",
			},
			"a",
		},
		{
			"right",
			args{
				"abcdefa",
			},
			"abcdefa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
