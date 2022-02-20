// Copyright 2022 <yaoyinnan@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func NewRandomList(v [][]int) *Node {
	return AddRandom(NewListNode(v), v)
}

func NewListNode(v [][]int) *Node {
	if len(v) == 0 {
		return nil
	}

	return &Node{
		Val:  v[0][0],
		Next: NewListNode(v[1:]),
	}
}

func AddRandom(h *Node, v [][]int) *Node {
	p := h
	var q *Node
	var num = 0
	for p != nil {
		if v[num][1] == -10001 {
			p.Random = nil
		} else {
			q = h
			for i := 0; i < v[num][1]; i++ {
				q = q.Next
			}
			p.Random = q
		}

		p = p.Next
		num += 1
	}

	return h
}

func (p *Node) Slice() []int {
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

func (p *Node) String() string {
	return fmt.Sprintf("%v", p.Slice())
}

func (p *Node) Array() [][]int {
	if p == nil {
		return [][]int{}
	}

	var arr [][]int
	q := p
	var s *Node
	for q != nil {
		num := 0
		if q.Random == nil {
			num = -10001
		} else {
			s = p
			for s != q.Random {
				s = s.Next
				num += 1
			}
		}
		arr = append(arr, []int{q.Val, num})
		q = q.Next
	}

	return arr
}
