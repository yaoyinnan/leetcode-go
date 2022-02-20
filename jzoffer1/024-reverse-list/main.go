// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

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

	fmt.Println(reverseList0(l))
}

func reverseList0(head *ListNode) *ListNode {
	p := head
	var q *ListNode
	var newHead *ListNode
	for p != nil {
		q = p
		p = p.Next
		q.Next = newHead
		newHead = q
	}

	return newHead
}
