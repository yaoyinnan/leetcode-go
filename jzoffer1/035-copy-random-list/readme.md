# 035. 复杂链表的复制

请实现 `copyRandomList` 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 `next` 指针指向下一个节点，还有一个 `random` 指针指向链表中的任意节点或者 `null`。

## 示例

示例 1：

![e1.png](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e1.png)

1. 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
2. 输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

示例 2：

![e2.png](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e2.png)

2. 输入：head = [[1,1],[2,1]]
3. 输出：[[1,1],[2,1]]

示例 3：

![e3.png](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e3.png)

1. 输入：head = [[3,null],[3,0],[3,null]]
2. 输出：[[3,null],[3,0],[3,null]]

示例 4：

1. 输入：head = []
2. 输出：[]

   解释：给定的链表为空（空指针），因此返回 null。

## 方案V0：双循环

### 思路及算法

双循环复制

### 代码

```go
func copyRandomList0(head *Node) *Node {
	var newHead = &Node{}
	p := head
	q := newHead

	for p != nil {
		item := Node{Val: p.Val}
		q.Next = &item
		q = q.Next
		p = p.Next
	}
	newHead = newHead.Next

	p = head
	s := head
	q = newHead
	t := newHead
	num := 0
	for p != nil {
		for s != p.Random {
			s = s.Next
			num += 1
		}
		for i := 0; i < num; i++ {
			t = t.Next
		}
		q.Random = t
		p = p.Next
		q = q.Next
		num = 0
		s = head
		t = newHead
	}

	return newHead
}
```

### 复杂度分析

- 时间复杂度：O(n^2)，其中 n 是链表的长度。对于每个节点，我们嵌套循环访问其「后继节点」及「随机指针指向的节点」。
- 空间复杂度：O(1)。

## 方案V1：回溯 + 哈希表

### 思路及算法

本题要求我们对一个特殊的链表进行深拷贝。如果是普通链表，我们可以直接按照遍历的顺序创建链表节点。而本题中因为随机指针的存在，当我们拷贝节点时，「当前节点的随机指针指向的节点」可能还没创建，因此我们需要变换思路。一个可行方案是，我们利用回溯的方式，让每个节点的拷贝操作相互独立。对于当前节点，我们首先要进行拷贝，然后我们进行「当前节点的后继节点」和「当前节点的随机指针指向的节点」拷贝，拷贝完成后将创建的新节点的指针返回，即可完成当前节点的两指针的赋值。

具体地，我们用哈希表记录每一个节点对应新节点的创建情况。遍历该链表的过程中，我们检查「当前节点的后继节点」和「当前节点的随机指针指向的节点」的创建情况。如果这两个节点中的任何一个节点的新节点没有被创建，我们都立刻递归地进行创建。当我们拷贝完成，回溯到当前层时，我们即可完成当前节点的指针赋值。注意一个节点可能被多个其他节点指向，因此我们可能递归地多次尝试拷贝某个节点，为了防止重复拷贝，我们需要首先检查当前节点是否被拷贝过，如果已经拷贝过，我们可以直接从哈希表中取出拷贝后的节点的指针并返回即可。

在实际代码中，我们需要特别判断给定节点为空节点的情况。

### 代码

```go
var cacheNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}

	if n, has := cacheNode[node]; has {
		return n
	}

	newNode := &Node{Val: node.Val}
	cacheNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList1(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return deepCopy(head)
}
```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。对于每个节点，我们至多访问其「后继节点」和「随机指针指向的节点」各一次，均摊每个点至多被访问两次。
- 空间复杂度：O(n)，其中 n 是链表的长度。为哈希表的空间开销。

## 方案V2：迭代 + 节点拆分

### 思路及算法

注意到方法一需要使用哈希表记录每一个节点对应新节点的创建情况，而我们可以使用一个小技巧来省去哈希表的空间。

我们首先将该链表中每一个节点拆分为两个相连的节点，例如对于链表 A→B→C，我们可以将其拆分为 A → A' → B → B' → C → C' 对于任意一个原节点 S，其拷贝节点 S' 即为其后继节点。

这样，我们可以直接找到每一个拷贝节点 S' 的随机指针应当指向的节点，即为其原节点 S 的随机指针指向的节点 T 的后继节点 T'。需要注意原节点的随机指针可能为空，我们需要特别判断这种情况。

当我们完成了拷贝节点的随机指针的赋值，我们只需要将这个链表按照原节点与拷贝节点的种类进行拆分即可，只需要遍历一次。同样需要注意最后一个拷贝节点的后继节点为空，我们需要特别判断这种情况。

### 代码

```go
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}
```

### 复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。我们只需要遍历该链表三次。
  - 读者们也可以自行尝试在计算拷贝节点的随机指针的同时计算其后继指针，这样只需要遍历两次。
- 空间复杂度：O(1)。注意返回值不计入空间复杂度。
