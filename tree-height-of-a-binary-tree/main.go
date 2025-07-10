package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	// Enter your code here. Read input from STDIN. Print output to STDOUT

	var (
		n, data int
		root    *Node
	)

	fmt.Scan(n)
	for i := 0; i < n; i++ {
		fmt.Scan(data)
		root = insert(root, data)
	}

	print(root)

	height := getHeight(root)
	fmt.Println(height)
}

func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Val: val}
	}

	if root.Val <= val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}

	return root
}

func print(root *Node) {
	if root == nil {
		return
	}

	print(root.Left)
	fmt.Printf("%d ", root.Val)
	print(root.Right)
}

func getHeight(root *Node) int {
	if root == nil {
		return 0
	}

	left := getHeight(root.Left)
	right := getHeight(root.Right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
