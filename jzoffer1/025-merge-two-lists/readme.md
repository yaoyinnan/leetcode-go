# 025. 合并两个排序的链表

输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

## 示例

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.

## 方案V0

### 思路及算法

先判断哪个链表第一个节点更小，然后进行去节点插入。

### 代码

```go
func mergeTwoLists0(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var low, high *ListNode
	if l1.Val < l2.Val {
		low, high = l1, l2
	} else {
		low, high = l2, l1
	}

	p, q := low, high
	var s, t *ListNode
	for p.Next != nil && q != nil {
		if q.Val <= p.Next.Val {
			s = q
			q = q.Next
			t = p.Next
			p.Next = s
			p.Next.Next = t
		} else {
			p = p.Next
		}
	}

	if q != nil {
		p.Next = q
	}

	return low
}
```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。
- 空间复杂度：O(1)。

## 方案V1：构造伪头节点

### 思路及算法

构造头节点，直接比对，将较小节点添加到后面。

### 代码

```go
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	dum := cur

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next, l1 = l1, l1.Next
		} else {
			cur.Next, l2 = l2, l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dum.Next
}
```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。
- 空间复杂度：O(1)。

## 方案V2：递归法

### 思路及算法

递归构建合并后的链表。

### 代码

```go
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}

```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。
- 空间复杂度：O(1)。