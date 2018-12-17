package addTwoNumbers

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。
	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		if l2.Val > 9 {
			if l2.Next != nil {
				l2.Next.Val = l2.Next.Val + 1
			} else {
				l2.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
			}
			return &ListNode{
				Val:  l2.Val % 10,
				Next: AddTwoNumbers(nil, l2.Next),
			}
		}
		return l2
	} else if l1 != nil && l2 == nil {
		if l1.Val > 9 {
			if l1.Next != nil {
				l1.Next.Val = l1.Next.Val + 1
			} else {
				l1.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
			}
			return &ListNode{
				Val:  l1.Val % 10,
				Next: AddTwoNumbers(nil, l1.Next),
			}
		}
		return l1
	} else if l1 == nil && l2 == nil {
		return nil
	}
	if l1.Val+l2.Val > 9 {
		if l1.Next != nil && l2.Next == nil {
			l1.Next.Val = l1.Next.Val + 1
			return &ListNode{
				Val:  (l1.Val + l2.Val) % 10,
				Next: AddTwoNumbers(l1.Next, nil),
			}
		} else if l2.Next != nil && l1.Next == nil {
			l2.Next.Val = l2.Next.Val + 1
			return &ListNode{
				Val:  (l1.Val + l2.Val) % 10,
				Next: AddTwoNumbers(nil, l2.Next),
			}
		} else if l1.Next == nil && l2.Next == nil {
			return &ListNode{
				Val: (l1.Val + l2.Val) % 10,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			}
		} else if l1.Next != nil && l2.Next != nil {
			l1.Next.Val = l1.Next.Val + 1
			return &ListNode{
				Val:  (l1.Val + l2.Val) % 10,
				Next: AddTwoNumbers(l1.Next, l2.Next),
			}
		}
	}

	return &ListNode{
		Val:  l1.Val + l2.Val,
		Next: AddTwoNumbers(l1.Next, l2.Next),
	}
}

func PrintNode(ln *ListNode) {
	for idx := ln; idx != nil; idx = idx.Next {
		fmt.Printf("%#v", idx)
	}
}
