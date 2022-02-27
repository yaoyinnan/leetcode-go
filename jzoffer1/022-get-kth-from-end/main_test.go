// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	for i, testCase := range tests {
		result := deleteNode(
			NewListNode(testCase.a...), testCase.v,
		)
		Assertf(t, reflect.DeepEqual(result, testCase.re),
			"%d: expect = %v, got = %v",
			i, testCase.re, result,
		)
	}
}

var tests = []struct {
	a, re []int
	v     int
}{
	{
		a:  []int{2, 4, 3},
		re: []int{4, 3},
		v:  2,
	},
	{
		a:  []int{2, 4, 3},
		re: []int{2, 3},
		v:  4,
	},
	{
		a:  []int{},
		re: []int{},
		v:  1,
	},
}
