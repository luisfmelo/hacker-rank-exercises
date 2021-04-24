package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeWithLevel struct {
	TreeNode
	Level int
}

type NodeQueue struct {
	q []TreeNodeWithLevel
}

func (q *NodeQueue) Push(n TreeNodeWithLevel) {
	q.q = append(q.q, n)
}

func (q *NodeQueue) Pop() TreeNodeWithLevel {
	if q.Empty() {
		panic("queue empty")
	}
	n := q.q[0]
	q.q = q.q[1:]
	return n
}

func (q *NodeQueue) Empty() bool {
	return len(q.q) == 0
}

func averageOfLevels(root *TreeNode) []float64 {
	counters := []float64{0}
	averages := []float64{0}

	q := &NodeQueue{}
	q.Push(TreeNodeWithLevel{*root, 0})

	for !q.Empty() {
		node := q.Pop()
		if len(counters) < node.Level+1 {
			counters = append(counters, 0)
			averages = append(averages, 0)
		}

		averages[node.Level] = (averages[node.Level]*counters[node.Level] + float64(node.Val)) / (counters[node.Level] + 1)
		counters[node.Level]++

		if node.Left != nil {
			q.Push(TreeNodeWithLevel{*node.Left, node.Level + 1})
		}
		if node.Right != nil {
			q.Push(TreeNodeWithLevel{*node.Right, node.Level + 1})
		}
	}
	return averages
}

func main() {
	r := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(averageOfLevels(r))
}
