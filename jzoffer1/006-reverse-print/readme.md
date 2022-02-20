# 006. 从尾到头打印链表

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

## 示例

1. 输入：`head = [1,3,2]`
2. 输出：`[2,3,1]`

## 准备工作

为了便于测试，为链表实现`NewListNode`构造函数和`Slice`方法。

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
```

## 方案V0

递归法

```go
func ReversePrint(head *ListNode) []int {
    if head == nil {
        return nil
    }
    return append(ReversePrint(head.Next), head.Val)
}
```

## 方案V1

非递归法

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