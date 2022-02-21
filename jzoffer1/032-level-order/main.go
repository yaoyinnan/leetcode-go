// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	root := []int{3, 9, 20, -1, -1, 15, 7}
	t := NewTree(root)

	fmt.Println(levelOrder(t))
}

func levelOrder(root *TreeNode) []int {
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
