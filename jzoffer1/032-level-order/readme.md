# 032 - I. 从上到下打印二叉树

从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

## 示例

示例：

例如:
给定二叉树:

```
[3,9,20,null,null,15,7],
```

    3
   / \
  9  20
    /  \
   15   7

返回：

```
[3,9,20,15,7]
```


## 方案V0

### 思路及算法

层次序列遍历二叉树

### 代码

```go
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var tempStack []*TreeNode
	var result []int

	p := root
	tempStack = append(tempStack, p)

	for len(tempStack) != 0 {
		p = tempStack[0]
		tempStack = tempStack[1:]
		result = append(result, p.Val)
		if p.Left != nil {
			tempStack = append(tempStack, p.Left)
		}
		if p.Right != nil {
			tempStack = append(tempStack, p.Right)
		}

	}

	return result
}
```

### 复杂度分析

- 时间复杂度 O(N) ：N 为二叉树的节点数量，即 BFS 需循环 N 次。
- 空间复杂度 O(N) ：最差情况下，即当树为平衡二叉树时，最多有 N/2 个树节点同时在 queue 中，使用 O(N) 大小的额外空间。