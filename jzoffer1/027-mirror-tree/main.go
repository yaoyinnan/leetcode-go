// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	root := []int{4, 2, 7, 1, 3, 6, 9}
	t := NewTree(root)

	fmt.Println(Print(mirrorTree(t)))
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var tempStack []*TreeNode

	p := root
	tempStack = append(tempStack, p)

	for len(tempStack) != 0 {
		p = tempStack[0]
		tempStack = tempStack[1:]
		if p.Left != nil {
			tempStack = append(tempStack, p.Left)
		}
		if p.Right != nil {
			tempStack = append(tempStack, p.Right)
		}
		p.Left, p.Right = p.Right, p.Left
	}

	return root
}
