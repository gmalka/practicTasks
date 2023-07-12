package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}

	for _, v := range delNodes(tree, []int{3, 5}) {
		fmt.Println(v.Val)
	}
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	m := make(map[int]interface{}, len(to_delete))

	for _, v := range to_delete {
		m[v] = nil
	}


	return recursive(root, nil, m)
}

func recursive(root, pre *TreeNode, m map[int]interface{}) []*TreeNode {
	if root == nil {
		return nil
	}

	var result []*TreeNode
	_, ok := m[root.Val]
	result = make([]*TreeNode, 0)

	if !ok {
		if pre == nil {
			result = append(result, root)
		}
		pre = root
	} else {
		if pre != nil {
			if pre.Left != nil && pre.Left.Val == root.Val {
				pre.Left = nil
			} else {
				pre.Right = nil
			}
		}
		pre = nil
	}

	result = append(result, recursive(root.Left, pre, m)...)
	result = append(result, recursive(root.Right, pre, m)...)

	return result
}