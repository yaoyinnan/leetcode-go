// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"testing"
)

func TestMirrorTree(t *testing.T) {
	for i, testCase := range tests {
		result := mirrorTree(
			NewTree(testCase.a),
		)
		Assertf(t, reflect.DeepEqual(result, testCase.c),
			"%d: expect = %v, got = %v",
			i, testCase.c, Print(result),
		)
	}
}

var tests = []struct {
	a, c []int
}{
	{
		a: []int{3, 9, 20, -1, -1, 15, 7},
		c: []int{3, 20, 9, 7, 15},
	},
	{
		a: []int{4, 2, 7, 1, 3, 6, 9},
		c: []int{4, 7, 2, 9, 6, 3, 1},
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
