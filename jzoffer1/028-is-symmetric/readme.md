# 028. 对称的二叉树

请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

## 示例

例如，二叉树[1,2,2,3,4,4,3] 是对称的。

```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

```
    1
   / \
  2   2
   \   \
   3    3
```

示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false

## 方案V0：递归比对

### 思路及算法

检查一棵二叉树是否对称，要检查以下几点：

1. 节点的左右节点相等，p.Left.Val == p.Right.Val
2. 节点左节点的左节点等于右节点的右节点，p.Left.Left.Val == p.Right.Right.Val
3. 节点左节点的右节点等于右节点的左节点，p.Left.Right.Val == p.Right.Left.Val

### 代码

```go
func recur(L *TreeNode, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	if L == nil || R == nil || L.Val != R.Val{
		return false
	}

	return recur(L.Left, R.Right) && recur(L.Right, R.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return recur(root.Left, root.Right)
	}
}
```

### 复杂度分析

- 时间复杂度 O(N) ： 其中 N 为二叉树的节点数量，每次执行 recur() 可以判断一对节点是否对称，因此最多调用 N/2 次 recur() 方法。
- 空间复杂度 O(N) ： 最差情况下（见下图），二叉树退化为链表，系统使用 O(N) 大小的栈空间。
