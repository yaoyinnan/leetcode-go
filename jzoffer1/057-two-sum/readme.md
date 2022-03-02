# 057. 和为s的两个数字

输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

## 示例

示例 1：

输入：nums = [2,7,11,15], target = 9

输出：[2,7] 或者 [7,2]


示例 2：

输入：nums = [10,26,30,31,47,60], target = 40

输出：[10,30] 或者 [30,10]


## 方案V0：双指针

### 思路及算法

双指针。

### 代码

```go
func twoSum(nums []int, target int) []int {
    i, j := 0, len(nums) - 1
    if i > j {
        return []int{}
    }
    sum := -1
    for true {
        sum = nums[i] + nums[j]
        if sum < target {
            i++
        } else if  sum > target {
            j--
        } else {
            return []int{nums[i], nums[j]}
        }
    }
    return []int{}
}
```

### 复杂度分析

- 时间复杂度：O(n)。
- 空间复杂度：O(1)。