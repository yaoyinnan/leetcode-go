# 024. 反转链表

定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

## 示例

1. 输入: 1->2->3->4->5->NULL
2. 输出: 5->4->3->2->1->NULL

## 准备工作

为了便于测试，为链表实现`NewListNode`构造函数、`Slice`方法和`Array`方法。

```go
func NewListNode(v ...int) *ListNode {
    if len(v) == 0 {
        return nil
    }
    return &ListNode{
        Val:  v[0],
        Next: NewListNode(v[1:]...),
    }
}

func (p *ListNode) Slice() []int {
    if p == nil {
        return nil
    }
    if p.Next == nil {
        return []int{p.Val}
    }
    return append([]int{p.Val},
        p.Next.Slice()...,
    )
}

func (p *ListNode) Array() []int {
	var arr []int
	q := p

	for q != nil {
		arr = append(arr, q.Val)
		q = q.Next
	}

return arr
}

```

## 方案V0

### 思路及算法

使用头插法构建翻转链表

### 代码

```go
func reverseList0(head *ListNode) *ListNode {
	p := head
	var q *ListNode
	var newHead *ListNode
	for p != nil{
		q = p
		p = p.Next
		q.Next = newHead
		newHead = q
	}

	return newHead
}
```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。

- 空间复杂度：O(1)。