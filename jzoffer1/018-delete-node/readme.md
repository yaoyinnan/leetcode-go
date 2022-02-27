# 018. 删除链表的节点

给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

## 示例

示例 1:

输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

## 方案V0

### 思路及算法

遍历，删除

### 代码

```go
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	p, q:= head.Next, head

	for p != nil {
		if p.Val == val {
			q.Next = p.Next
		}
		p = p.Next
		q = q.Next
	}

	return head
}
```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。
- 空间复杂度：O(1)。