package addStrings

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "good",
			args: args{
				"123",
				"123",
			},
			want: "246",
		},
		{
			name: "good",
			args: args{
				"600",
				"500",
			},
			want: "1100",
		},
		{
			name: "good",
			args: args{
				"1",
				"1",
			},
			want: "2",
		},
		{
			name: "good",
			args: args{
				"0",
				"0",
			},
			want: "0",
		},
		{
			name: "good",
			args: args{
				"999",
				"1",
			},
			want: "1000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
