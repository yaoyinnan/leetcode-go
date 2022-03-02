# 021. 调整数组顺序使奇数位于偶数前面

输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。

## 示例

示例：

输入：nums = [1,2,3,4]

输出：[1,3,2,4]

注：[3,1,2,4] 也是正确的答案之一。

## 方案V0：双指针

### 思路及算法

双指针或者双端队列。

### 代码

```go
func exchange(nums []int) []int {
	i, j := 0, len(nums) - 1
	for i < j {
		if nums[i] % 2 == 1 {
			i++
			continue
		}
		if nums[j] % 2 == 0 {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
```

### 复杂度分析

- 时间复杂度：O(n)。
- 空间复杂度：O(1)。