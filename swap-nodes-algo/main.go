package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'swapNodes' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY indexes
 *  2. INTEGER_ARRAY queries
 */

type Node struct {
	Val   int32
	Left  *Node
	Right *Node
}

// Queue represents slice of any generic type
type Queue[T any] []T

// Enqueue append data from back of queue
func (q *Queue[T]) Enqueue(val T) {
	*q = append(*q, val)
}

// Dequeue pop data from front of queue
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(*q) == 0 {
		return zero, false
	}
	val := (*q)[0]
	(*q)[0] = zero
	*q = (*q)[1:]
	return val, true
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	res := make([][]int32, len(queries))

	root := build(indexes)
	// tmp := []int32{}
	// traverse(root, &tmp)
	// fmt.Println()

	for i, k := range queries {
		swap(root, k, 1)
		traverse(root, &res[i])
		// fmt.Println()
	}

	return res
}

// build a tree from `indexes` of 2D-array on level order
// return a root of the tree
func build(indexes [][]int32) *Node {
	root := &Node{Val: 1}
	var queue Queue[*Node]
	queue.Enqueue(root)

	for i := 0; i < len(indexes); i++ {
		curr, ok := queue.Dequeue()
		if !ok {
			break
		}

		if indexes[i][0] > -1 {
			curr.Left = &Node{Val: indexes[i][0]}
			queue.Enqueue(curr.Left)
		}

		if indexes[i][1] > -1 {
			curr.Right = &Node{Val: indexes[i][1]}
			queue.Enqueue(curr.Right)
		}
	}

	return root
}

// traverse transform a tree into in-order structure from a root
// update in arr
func traverse(root *Node, arr *[]int32) {
	if root == nil {
		return
	}

	traverse(root.Left, arr)
	*arr = append(*arr, root.Val)
	// fmt.Printf("%d ", root.Val)
	traverse(root.Right, arr)
}

// swap all nodes in h depths (which is multiplication of k) of tree from a root
func swap(root *Node, k, l int32) {
	if root == nil {
		return
	}

	if l%k == 0 {
		root.Left, root.Right = root.Right, root.Left
	}

	swap(root.Left, k, l+1)
	swap(root.Right, k, l+1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var indexes [][]int32
	for i := 0; i < int(n); i++ {
		indexesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != 2 {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for i, rowItem := range result {
		for j, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if j != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
