# 032 - III. 从上到下打印二叉树 III

请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

## 示例

示例：

例如:
给定二叉树:

```
[3,9,20,null,null,15,7],
```

```
    3
   / \
  9  20
    /  \
   15   7
```

返回其层次遍历结果：


```
[
    [3],
    [20,9],
    [15,7]
]
```

## 方案V0：层序遍历 + 双端队列

### 思路及算法

- 利用双端队列的两端皆可添加元素的特性，设打印列表（双端队列） tmp ，并规定： 
  - 奇数层 则添加至 tmp 尾部 ， 
  - 偶数层 则添加至 tmp 头部 。
算法流程：
1. 特例处理： 当树的根节点为空，则直接返回空列表 [] ；
2. 初始化： 打印结果空列表 res ，包含根节点的双端队列 deque ；
3. BFS 循环： 当 deque 为空时跳出； 
   1. 新建列表 tmp ，用于临时存储当前层打印结果； 
   2. 当前层打印循环： 循环次数为当前层节点数（即 deque 长度）； 
      1. 出队： 队首元素出队，记为 node； 
      2. 打印： 若为奇数层，将 node.val 添加至 tmp 尾部；否则，添加至 tmp 头部； 
      3. 添加子节点： 若 node 的左（右）子节点不为空，则加入 deque ； 
   3. 将当前层结果 tmp 转化为 list 并添加入 res ；
4. 返回值： 返回打印结果列表 res 即可；

### 代码

```go
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var Q []*TreeNode

	p := root
	Q = append(Q, p)

	for len(Q) != 0 {
		var tmp []int

		for i := len(Q); i > 0; i-- {
			p = Q[0]
			Q = Q[1:]

			if len(result)%2 == 0 {
				tmp = append(tmp, p.Val)
			} else {
				tmp = append([]int{p.Val}, tmp...)
			}

			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
		}

		result = append(result, tmp)
	}

	return result
}
```

### 复杂度分析

- 时间复杂度 O(N) ：N 为二叉树的节点数量，即 BFS 需循环 N 次。
- 空间复杂度 O(N) ：最差情况下，即当树为平衡二叉树时，最多有 N/2 个树节点同时在 queue 中，使用 O(N) 大小的额外空间。