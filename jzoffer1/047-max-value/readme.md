# 047. 礼物的最大价值

在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

## 示例

输入:
```bigquery
[
    [1,3,1],
    [1,5,1],
    [4,2,1]
]
```
输出: 12

解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

## 方案V0：动态规划

### 思路及算法

遍历二维数组，当前结点值加上上结点和左节点的最大值，最后一个值即为最大礼物价值。
其中，第 0 行和第 0 列需要单独判断。

### 代码

```go
func maxValue0(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var max int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i - 1 < 0 && j - 1 < 0{
				continue
			} else if i - 1 >= 0 && j - 1 >= 0 {
				if grid[i - 1][j] < grid[i][j - 1] {
					max = grid[i][j - 1]
				} else {
					max = grid[i - 1][j]
				}
				grid[i][j] += max
			} else if i - 1 < 0 {
				grid[i][j] += grid[i][j - 1]
			} else if j - 1 < 0 {
				grid[i][j] += grid[i - 1][j]
			}
		}
	}
	return grid[m - 1][n - 1]
}
```

对于第 0 行和第 0 列，可以单独判断以节省在嵌套循环中的浪费。

```go
func maxValue1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i - 1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j - 1]
	}
	var max int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i - 1 >= 0 && j - 1 >= 0 {
				if grid[i - 1][j] < grid[i][j - 1] {
					max = grid[i][j - 1]
				} else {
					max = grid[i - 1][j]
				}
				grid[i][j] += max
			}
		}
	}
	return grid[m - 1][n - 1]
}
```

### 复杂度分析

- 时间复杂度 O(MN) ： M, N 分别为矩阵行高、列宽；动态规划需遍历整个 grid 矩阵，使用 O(MN) 时间。
- 空间复杂度 O(1) ： 原地修改使用常数大小的额外空间。
