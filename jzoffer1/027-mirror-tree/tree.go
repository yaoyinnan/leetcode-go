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
			if v[i] != -1 {
				temp.Left = &TreeNode{Val: v[i]}
				queue = append(queue, temp.Left)
			}
			i++
		}
		if i < len(v) {
			if v[i] != -1 {
				temp.Right = &TreeNode{Val: v[i]}
				queue = append(queue, temp.Right)
			}
			i++
		}
	}

	return root
}

func Print(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var tempStack []*TreeNode
	var result []int

	p := root
	tempStack = append(tempStack, p)

	for len(tempStack) != 0 {
		p = tempStack[0]
		tempStack = tempStack[1:]
		result = append(result, p.Val)
		if p.Left != nil {
			tempStack = append(tempStack, p.Left)
		}
		if p.Right != nil {
			tempStack = append(tempStack, p.Right)
		}
	}

	return result
}
