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
		result := reverseList0(
			NewListNode(testCase.a...),
		)
		Assertf(t, reflect.DeepEqual(result.Array(), testCase.re),
			"%d: expect = %v, got = %v",
			i, testCase.re, result.Array(),
		)
	}
}

var tests = []struct {
	a, re []int
}{
	{
		a:  []int{2, 4, 3},
		re: []int{3, 4, 2},
	},
	{
		a:  []int{7, 3},
		re: []int{3, 7},
	},
	{
		a:  []int{5},
		re: []int{5},
	},
}
