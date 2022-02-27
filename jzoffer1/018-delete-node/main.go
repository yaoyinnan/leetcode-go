// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	l := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	fmt.Println(deleteNode(l, 2))
	fmt.Println(deleteNode(l, 3))
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	p, q := head.Next, head

	for p != nil {
		if p.Val == val {
			q.Next = p.Next
		}
		p = p.Next
		q = q.Next
	}

	return head
}
