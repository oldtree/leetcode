package longestSubstringWithoutRepeatingCharacters

import "testing"

func Test_LengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"right",
			args{
				"abcabcbb",
			},
			3,
		},

		{
			"right",
			args{
				"bbbbbb",
			},
			1,
		},
		{
			"right",
			args{
				"abababababab",
			},
			2,
		},
		{
			"right",
			args{
				"aaaaabbbbbbb",
			},
			2,
		},
		{
			"right",
			args{
				"a",
			},
			1,
		},
		{
			"right",
			args{
				"aabcdefg",
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
