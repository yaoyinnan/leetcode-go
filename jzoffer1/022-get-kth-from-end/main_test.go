// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"testing"
)

func TestGetKthFromEnd(t *testing.T) {
	for i, testCase := range tests {
		result := getKthFromEnd(
			NewListNode(testCase.a...), testCase.k,
		)
		Assertf(t, reflect.DeepEqual(result, testCase.re),
			"%d: expect = %v, got = %v",
			i, testCase.re, result,
		)
	}
}

var tests = []struct {
	a, re []int
	k     int
}{
	{
		a:  []int{2, 4, 3},
		re: []int{4, 3},
		k:  2,
	},
	{
		a:  []int{2, 4, 3},
		re: []int{3},
		k:  1,
	},
}
