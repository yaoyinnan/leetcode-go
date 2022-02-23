// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(v []int) *TreeNode {
	if len(v) == 0 {
		return nil
	}
	var queue []*TreeNode

	i := 0
	root := &TreeNode{Val: v[i]}
	i++
	queue = append(queue, root)

	for len(queue) != 0 {
		temp := queue[0]
		queue = queue[1:]
		if i < len(v) {
			if v[i] != -1000 {
				temp.Left = &TreeNode{Val: v[i]}
				queue = append(queue, temp.Left)
			}
			i++
		}
		if i < len(v) {
			if v[i] != -1000 {
				temp.Right = &TreeNode{Val: v[i]}
				queue = append(queue, temp.Right)
			}
			i++
		}
	}

	return root
}
