// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	fmt.Println(maxValue0([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(maxValue1([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func maxValue0(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var max int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 < 0 && j-1 < 0 {
				continue
			} else if i-1 >= 0 && j-1 >= 0 {
				if grid[i-1][j] < grid[i][j-1] {
					max = grid[i][j-1]
				} else {
					max = grid[i-1][j]
				}
				grid[i][j] += max
			} else if i-1 < 0 {
				grid[i][j] += grid[i][j-1]
			} else if j-1 < 0 {
				grid[i][j] += grid[i-1][j]
			}
		}
	}
	return grid[m-1][n-1]
}

func maxValue1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	var max int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				if grid[i-1][j] < grid[i][j-1] {
					max = grid[i][j-1]
				} else {
					max = grid[i-1][j]
				}
				grid[i][j] += max
			}
		}
	}
	return grid[m-1][n-1]
}
