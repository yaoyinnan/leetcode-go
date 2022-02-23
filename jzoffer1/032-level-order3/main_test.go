// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	for i, testCase := range tests {
		result := levelOrder(
			NewTree(testCase.a),
		)
		Assertf(t, reflect.DeepEqual(result, testCase.c),
			"%d: expect = %v, got = %v",
			i, testCase.c, result,
		)
	}
}

var tests = []struct {
	a []int
	c [][]int
}{
	{
		a: []int{3, 9, 20, -1000, -1000, 15, 7},
		c: [][]int{{3}, {20, 9}, {15, 7}},
	},
	{
		a: []int{0, 2, 4, 1, -1000, 3, -1, 5, 1, -1000, 6, -1000, 8},
		c: [][]int{{0}, {4, 2}, {1, 3, -1}, {8, 6, 1, 5}},
	},
	{
		a: []int{1, -1000, 2, -1000, 3},
		c: [][]int{{1}, {2}, {3}},
	},
	{
		a: []int{1, 2, -1000, 3, -1000, -1000, -1000},
		c: [][]int{{1}, {2}, {3}},
	},
	{
		a: []int{},
		c: [][]int{},
	},
}
