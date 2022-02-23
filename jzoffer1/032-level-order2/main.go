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

	fmt.Println(levelOrder0(t))
	fmt.Println(levelOrder1(t))
}

func levelOrder0(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var Q []*TreeNode
	var layer []int
	currentLayer := 0
	p := root

	Q = append(Q, p)
	layer = append(layer, currentLayer)
	result = append(result, []int{})

	for len(Q) != 0 {
		p = Q[0]
		Q = Q[1:]
		currentLayer = layer[0]
		layer = layer[1:]
		result[currentLayer] = append(result[currentLayer], p.Val)
		if p.Left != nil {
			Q = append(Q, p.Left)
			layer = append(layer, currentLayer+1)
		}
		if p.Right != nil {
			Q = append(Q, p.Right)
			layer = append(layer, currentLayer+1)
		}
		if p.Left != nil || p.Right != nil {
			currentLayer++
			if len(result) <= currentLayer {
				result = append(result, []int{})
			}
		}
	}

	return result
}

func levelOrder1(root *TreeNode) [][]int {
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
			tmp = append(tmp, p.Val)

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
