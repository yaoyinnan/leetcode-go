// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists0(l1, l2))

	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists1(l1, l2))

	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists2(l1, l2))
}

func mergeTwoLists0(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var low, high *ListNode
	if l1.Val < l2.Val {
		low, high = l1, l2
	} else {
		low, high = l2, l1
	}

	p, q := low, high
	var s, t *ListNode
	for p.Next != nil && q != nil {
		if q.Val <= p.Next.Val {
			s = q
			q = q.Next
			t = p.Next
			p.Next = s
			p.Next.Next = t
		} else {
			p = p.Next
		}
	}

	if q != nil {
		p.Next = q
	}

	return low
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	dum := cur

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next, l1 = l1, l1.Next
		} else {
			cur.Next, l2 = l2, l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dum.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}
