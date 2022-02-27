// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(getKthFromEnd(l, 2))
	fmt.Println(getKthFromEnd(l, 4))
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p, q := head, head
	for i := 0; i < k; i++ {
		p = p.Next
	}
	for p != nil {
		p, q = p.Next, q.Next
	}

	return q
}
