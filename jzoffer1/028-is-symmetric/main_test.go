// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	for i, testCase := range tests {
		result := isSymmetric(
			NewTree(testCase.a),
		)
		Assertf(t, result == testCase.flag,
			"%d: expect = %v, got = %v",
			i, testCase.flag, result,
		)
	}
}

var tests = []struct {
	a    []int
	flag bool
}{
	{
		a:    []int{1, 2, 2, 3, 4, 4, 3},
		flag: true,
	},
	{
		a:    []int{1, 2, 2, -1000, 3, -1000, 3},
		flag: false,
	},
	{
		a:    []int{},
		flag: true,
	},
}
