// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	root := []int{0, 2, 4, 1, -1000, 3, -1, 5, 1, -1000, 6, -1000, 8}
	t := NewTree(root)

	fmt.Println(levelOrder(t))
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var Q []*TreeNode

	p := root
	Q = append(Q, p)

	for len(Q) != 0 {
		var tmp []int

		for i := len(Q); i > 0; i-- {
			p = Q[0]
			Q = Q[1:]

			if len(result)%2 == 0 {
				tmp = append(tmp, p.Val)
			} else {
				tmp = append([]int{p.Val}, tmp...)
			}

			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
		}

		result = append(result, tmp)
	}

	return result
}
