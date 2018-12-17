package addTwoNumbers

import (
	"reflect"
	"testing"
)

var l1 = &ListNode{
	1, nil,
}

var l2 = &ListNode{
	9, &ListNode{
		9, nil,
	},
}

var l3 = &ListNode{
	0, &ListNode{
		0, &ListNode{
			1, nil,
		},
	},
}

var l4 = &ListNode{
	2, &ListNode{
		4, &ListNode{
			3, nil,
		},
	},
}

var l5 = &ListNode{
	5, &ListNode{
		6, &ListNode{
			4, nil,
		},
	},
}

var l6 = &ListNode{
	7, &ListNode{
		0, &ListNode{
			8, nil,
		},
	},
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"right",
			args{
				nil,
				nil,
			},
			nil,
		},
		{
			"right",
			args{
				l1,
				l2,
			},
			l3,
		},
		{
			"right",
			args{
				l4,
				l5,
			},
			l6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
