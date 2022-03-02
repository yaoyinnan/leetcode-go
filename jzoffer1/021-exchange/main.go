// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))

}

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 1 {
			i++
			continue
		}
		if nums[j]%2 == 0 {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
