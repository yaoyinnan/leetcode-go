// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSubStructure(NewTree([]int{1, 2, 3}), NewTree([]int{3, 1})))
	fmt.Println(isSubStructure(NewTree([]int{3, 4, 5, 1, 2}), NewTree([]int{4, 1})))
}

func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return A != nil && B != nil &&
		(recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}
