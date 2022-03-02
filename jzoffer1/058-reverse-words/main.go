// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))

}

func reverseWords(s string) string {
	var res string
	for i, j := len(s)-1, len(s)-1; i >= -1 && j >= 0; i-- {
		if s[j] == ' ' {
			j--
		}
		if (i == -1 || s[i] == ' ') && i < j {
			res = strings.Join([]string{res, s[i+1 : j+1]}, " ")
			j = i
		}
	}
	if len(res) == 0 {
		return ""
	}
	return res[1:]
}
