# 052. 两个链表的第一个公共节点

输入两个链表，找出它们的第一个公共节点。

如下面的两个链表：

![img.jpg](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_statement.png)

在节点 c1 开始相交。


示例 1：

![img.jpg](https://assets.leetcode.com/uploads/2018/12/13/160_example_1.png)

输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3

输出：Reference of the node with value = 8

输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。


示例2：

![img.jpg](https://assets.leetcode.com/uploads/2018/12/13/160_example_2.png)

输入：intersectVal= 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1

输出：Reference of the node with value = 2

输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

示例3：

![img.jpg](https://assets.leetcode.com/uploads/2018/12/13/160_example_3.png)

输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2

输出：null

输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 null。

注意：

- 如果两个链表没有交点，返回 null.
- 在返回结果后，两个链表仍须保持原有的结构。
- 可假定整个链表结构中没有循环。
- 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
- 本题与主站 160 题相同：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

## 示例

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.

## 方案V0：哈希表

### 思路及算法

1. 先遍历HeadA，以节点为key存入哈希表，值为boll。
2. 再遍历HeadB，若在哈希表中，则返回。
3. 若遍历后没有公共节点，则最后返回nil。

### 代码

```go
func getIntersectionNode0(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for p := headA; p != nil; p = p.Next {
		vis[p] = true
	}
	for p := headB; p != nil; p = p.Next {
		if vis[p] {
			return p
		}
	}
	return nil
}
```

### 复杂度分析

- 时间复杂度：O(M+N)，遍历HeadA和HeadB。
- 空间复杂度：O(M)。

## 方案V1：双指针

### 思路及算法

1. 由于上述哈希表的方法需要建立哈希表，占用了M个空间，可以采用双指针进行优化，将空间复杂度降至 O(1)。
2. 方案
   1. 首先判断headA和headB是否其一为空，若为空则不存在公共节点，返回nil。
   2. 每步操作需要同时更新p1和p2。当遍历到链表尾部时，将对应指针指向另一个链表头。
   3. 上述操作，p1遍历了a+c+b个节点，p2遍历了b+c+a个节点，将同时到达第一个公共节点。

### 代码

```go
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}

	return p1
}
```

### 复杂度分析

- 时间复杂度：O(M+N)，遍历HeadA和HeadB。
- 空间复杂度：O(1)。
