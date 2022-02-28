// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	c := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}
	l1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  8,
				Next: c,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: c,
				},
			},
		},
	}

	l3 := getIntersectionNode0(l1, l2)
	if l3 != nil {
		fmt.Println(l3.Val)
	} else {
		fmt.Println("null")
	}

	l3 = getIntersectionNode1(l1, l2)
	if l3 != nil {
		fmt.Println(l3.Val)
	} else {
		fmt.Println("null")
	}
}

func getIntersectionNode0(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for p := headA; p != nil; p = p.Next {
		vis[p] = true
	}
	for p := headB; p != nil; p = p.Next {
		if vis[p] {
			return p
		}
	}
	return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}

	return p1
}
