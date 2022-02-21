// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"testing"
)

func TestReversePrint0(t *testing.T) {
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
	a, c []int
}{
	{
		a: []int{3, 9, 20, -1, -1, 15, 7},
		c: []int{3, 9, 20, 15, 7},
	},
	{
		a: []int{1, -1, 2, -1, 3},
		c: []int{1, 2, 3},
	},
	{
		a: []int{1, 2, -1, 3, -1, -1, -1},
		c: []int{1, 2, 3},
	},
	{
		a: []int{},
		c: []int{},
	},
}
