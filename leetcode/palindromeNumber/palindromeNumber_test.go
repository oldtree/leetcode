package palindromeNumber

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				x: 0,
			},
			want: true,
		},
		{
			name: "-1",
			args: args{
				x: -1,
			},
			want: false,
		},
		{
			name: "10",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "11",
			args: args{
				x: 11,
			},
			want: true,
		},
		{
			name: "101",
			args: args{
				x: 101,
			},
			want: true,
		},
		{
			name: "331111101111133",
			args: args{
				x: 331111101111133,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
