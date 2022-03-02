// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 5))

}

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	if i > j {
		return []int{}
	}
	sum := -1
	for true {
		sum = nums[i] + nums[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{}
}
