# 026. 树的子结构

输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

## 示例

示例：

例如:
给定的树 A:

```
     3
    / \
   4   5
  / \
 1   2
```

给定的树 B：


```
   4 
  /
 1
```

返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true

## 方案V0

### 思路及算法

以每个节点作为一棵子树，递归遍历每棵子树，与B进行依次对照。

### 代码

```go
func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return A != nil && B != nil &&
		(recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

```

### 复杂度分析

时间复杂度 O(M) ： 其中 M,N 分别为树 A 和 树 B 的节点数量；先序遍历树 A 占用 O(M) ，每次调用 recur(A, B) 判断占用 O(N) 。
空间复杂度 O(M) ： 当树 A 和树 B 都退化为链表时，递归调用深度最大。当 M ≤ N 时，遍历树 A 与递归判断的总递归深度为 M ；当 M > N 时，最差情况为遍历至树 A 叶子节点，此时总递归深度为 M。