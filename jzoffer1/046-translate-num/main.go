// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(translateNum(12258))

}

func translateNum(num int) int {
	str := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	var pre string
	for i := 0; i < len(str); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre = str[i-1 : i+1]
		if pre >= "10" && pre <= "25" {
			r += p
		}
	}
	return r
}
