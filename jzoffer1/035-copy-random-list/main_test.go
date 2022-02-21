// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"testing"
)

func TestCopyRandomList0(t *testing.T) {
	for i, testCase := range tests {
		result := copyRandomList0(
			NewRandomList(testCase.a),
		)
		Assertf(t, reflect.DeepEqual(result.Array(), testCase.c),
			"%d: expect = %v, got = %v",
			i, testCase.c, result.Array(),
		)
	}
}

func TestCopyRandomList1(t *testing.T) {
	for i, testCase := range tests {
		result := copyRandomList1(
			NewRandomList(testCase.a),
		)
		Assertf(t, reflect.DeepEqual(result.Array(), testCase.c),
			"%d: expect = %v, got = %v",
			i, testCase.c, result.Array(),
		)
	}
}

func TestCopyRandomList2(t *testing.T) {
	for i, testCase := range tests {
		result := copyRandomList2(
			NewRandomList(testCase.a),
		)
		Assertf(t, reflect.DeepEqual(result.Array(), testCase.c),
			"%d: expect = %v, got = %v",
			i, testCase.c, result.Array(),
		)
	}
}

var tests = []struct {
	a, c [][]int
}{
	{
		a: [][]int{{7, -10001}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		c: [][]int{{7, -10001}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
	},
	{
		a: [][]int{{1, 1}, {2, 1}},
		c: [][]int{{1, 1}, {2, 1}},
	},
	{
		a: [][]int{{3, -10001}, {3, 0}, {3, -10001}},
		c: [][]int{{3, -10001}, {3, 0}, {3, -10001}},
	},
	{
		a: [][]int{},
		c: [][]int{},
	},
}
