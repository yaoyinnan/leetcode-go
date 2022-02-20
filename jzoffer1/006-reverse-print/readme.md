# 006. 从尾到头打印链表

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

## 示例

1. 输入：`head = [1,3,2]`
2. 输出：`[2,3,1]`

## 方案V0

### 思路及算法

递归法

### 代码

```go
func ReversePrint(head *ListNode) []int {
    if head == nil {
        return nil
    }
    return append(ReversePrint(head.Next), head.Val)
}
```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。
- 空间复杂度：O(1)。

## 方案V1

### 思路及算法

非递归法

### 代码

```go
func reversePrint1(head *ListNode) []int {
	var result []int

	p := head

	for p != nil{
		result = append(result, p.Val)
		p = p.Next
	}

	var rLen = len(result)
	for i := 0; i < rLen / 2; i++{
		result[i], result[rLen - i - 1] = result[rLen - i - 1], result[i]
	}

	return result
}
```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。循环两次，先构建数组再翻转，平均为n。
- 空间复杂度：O(n)。