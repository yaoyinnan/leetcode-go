# 032 - II. 从上到下打印二叉树 II

从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

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
    [9,20],
    [15,7]
]
```

## 方案V0

### 思路及算法

I. 按层打印： 题目要求的二叉树的 从上至下 打印（即按层打印），又称为二叉树的 广度优先搜索（BFS）。BFS 通常借助 队列 的先入先出特性来实现。

II. 每层打印到一行： 将本层全部节点打印到一行，并将下一层全部节点加入队列，以此类推，即可分为多行打印。

算法流程：
1. 特例处理： 当根节点为空，则返回空列表 [] ；
2. 初始化： 打印结果列表 res = [] ，包含根节点的队列 queue = [root] ，包含节点所在层的队列 layer = [0]；
3. BFS 循环： 当队列 queue 为空时跳出；
   1. 出队： 队首元素和所在层出队，记为 node 和 currentLayer；
   2. 打印： 将 node.val 添加至对应层；
   3. 添加子节点： 若 node 的左（右）子节点不为空，则将左（右）子节点加入队列 queue ；
   4. 将结果添加到对应层。
4. 返回值： 返回打印结果列表 res 即可。

### 代码

```go
func levelOrder0(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var Q []*TreeNode
	var result [][]int
	var layer []int
	currentLayer := 0
	p := root

	Q = append(Q, p)
	layer = append(layer, currentLayer)
	result = append(result, []int{})

	for len(Q) != 0 {
		p = Q[0]
		Q = Q[1:]
		currentLayer = layer[0]
		layer = layer[1:]
		result[currentLayer] = append(result[currentLayer], p.Val)
		if p.Left != nil {
			Q = append(Q, p.Left)
			layer = append(layer, currentLayer + 1)
		}
		if p.Right != nil {
			Q = append(Q, p.Right)
			layer = append(layer, currentLayer + 1)
		}
		if p.Left != nil || p.Right != nil {
			currentLayer++
			if len(result) <= currentLayer{
				result = append(result, []int{})
			}
		}
	}

	return result
}
```

### 复杂度分析

- 时间复杂度 O(N) ：N 为二叉树的节点数量，即 BFS 需循环 N 次。
- 空间复杂度 O(N) ：最差情况下，即当树为平衡二叉树时，最多有 N/2 个树节点同时在 queue 中，使用 O(N) 大小的额外空间。


## 方案V1

### 思路及算法

I. 按层打印： 题目要求的二叉树的 从上至下 打印（即按层打印），又称为二叉树的 广度优先搜索（BFS）。BFS 通常借助 队列 的先入先出特性来实现。

II. 每层打印到一行： 将本层全部节点打印到一行，并将下一层全部节点加入队列，以此类推，即可分为多行打印。

算法流程：
1. 特例处理： 当根节点为空，则返回空列表 [] ；
2. 初始化： 打印结果列表 res = [] ，包含根节点的队列 queue = [root] ；
3. BFS 循环： 当队列 queue 为空时跳出；
   1. 新建一个临时列表 tmp ，用于存储当前层打印结果； 
   2. 当前层打印循环： 循环次数为当前层节点数（即队列 queue 长度）； 
      1. 出队： 队首元素出队，记为 node； 
      2. 打印： 将 node.val 添加至 tmp 尾部； 
      3. 添加子节点： 若 node 的左（右）子节点不为空，则将左（右）子节点加入队列 queue ； 
   3. 将当前层结果 tmp 添加入 res 。
4. 返回值： 返回打印结果列表 res 即可。

### 代码

```go
func levelOrder1(root *TreeNode) [][]int {
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
			tmp = append(tmp, p.Val)

			if p.Left != nil{
				Q = append(Q, p.Left)
			}
			if p.Right != nil{
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