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

	fmt.Println(reversePrint0(l))
	fmt.Println(reversePrint1(l))
}

func reversePrint0(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint0(head.Next), head.Val)
}

func reversePrint1(head *ListNode) []int {
	var result []int

	p := head

	for p != nil {
		result = append(result, p.Val)
		p = p.Next
	}

	var rLen = len(result)
	for i := 0; i < rLen/2; i++ {
		result[i], result[rLen-i-1] = result[rLen-i-1], result[i]
	}

	return result
}
