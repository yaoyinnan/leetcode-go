// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSymmetric(NewTree([]int{1, 2, 2, 3, 4, 4, 3})))
	fmt.Println(isSymmetric(NewTree([]int{1, 2, 2, -1000, 3, -1000, 3})))
}

func recur(L *TreeNode, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	if L == nil || R == nil || L.Val != R.Val {
		return false
	}

	return recur(L.Left, R.Right) && recur(L.Right, R.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return recur(root.Left, root.Right)
	}
}
