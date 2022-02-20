// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	head := [][]int{{7, -10001}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	l := NewRandomList(head)

	fmt.Println(copyRandomList0(l))
	fmt.Println(copyRandomList1(l))
	fmt.Println(copyRandomList2(l))
}

// V0：双循环拷贝
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

// V1：回溯 + 哈希表
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

// V2：迭代 + 节点拆分
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
