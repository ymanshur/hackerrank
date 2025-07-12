package main

import "fmt"

type TreeNode struct {
	Val   int32
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, v1, v2 int32) *TreeNode {
	if root.Val < v1 && root.Val < v2 {
		return traverse(root.Right, v1, v2)
	}

	if root.Val > v1 && root.Val > v2 {
		return traverse(root.Left, v1, v2)
	}

	return root
}

func insert(root *TreeNode, val int32) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val <= root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}

	return root
}

// v1 <= lcm <= v2

func main() {
	var (
		n, data, v1, v2 int32
		root            *TreeNode
	)

	fmt.Scan(&n)
	for range n {
		fmt.Scan(&data)
		root = insert(root, data)
	}

	fmt.Scan(&v1, &v2)

	ans := traverse(root, v1, v2)
	fmt.Println(ans.Val)
}
