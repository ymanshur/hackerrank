package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue []int32

func (q *Queue) Enqueue(val int32) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() (int32, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	val := (*q)[0]
	(*q)[0] = 0
	*q = (*q)[1:]
	return val, true
}

func (q *Queue) Len() int64 {
	return int64(len(*q))
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Complete the findShortest function below.

/*
 * For the unweighted graph, <name>:
 *
 * 1. The number of nodes is <name>Nodes.
 * 2. The number of edges is <name>Edges.
 * 3. An edge exists between <name>From[i] to <name>To[i].
 *
 */
func findShortest(graphNodes int32, graphFrom []int32, graphTo []int32, ids []int64, val int64) int32 {
	graph := make([][]int32, graphNodes)
	for i := 0; i < len(graphFrom); i++ {
		graph[graphFrom[i]-1] = append(graph[graphFrom[i]-1], graphTo[i]-1)
		graph[graphTo[i]-1] = append(graph[graphTo[i]-1], graphFrom[i]-1)
	}

	var (
		queue      Queue
		visited    = make([]bool, graphNodes)
		pathLength = make(map[int32]int32, graphNodes)
	)

	for i, color := range ids {
		if color == val {
			queue.Enqueue(int32(i))
			pathLength[int32(i)] = 0
		}
	}

	if queue.Len() < 2 {
		return -1
	}

	for {
		from, ok := queue.Dequeue()
		if !ok {
			break
		}

		visited[from] = true
		// fmt.Println("from:\t", from+1)

		for _, to := range graph[from] {
			// fmt.Println("to:\t", to+1)

			if val, ok := pathLength[to]; ok && !visited[to] {
				return val + pathLength[from] + 1
			}

			queue.Enqueue(to)
			pathLength[to] = pathLength[from] + 1
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20) // 1 MB buffer

	// Read graphNodes and graphEdges
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	graphNodes, _ := strconv.Atoi(parts[0])
	graphEdges, _ := strconv.Atoi(parts[1])

	var graphFrom, graphTo []int32
	for i := 0; i < graphEdges; i++ {
		line, _ = reader.ReadString('\n')
		parts = strings.Fields(line)
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		graphFrom = append(graphFrom, int32(from))
		graphTo = append(graphTo, int32(to))
	}

	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	ids := make([]int64, graphNodes)
	for i := 0; i < graphNodes; i++ {
		id, _ := strconv.ParseInt(parts[i], 10, 64)
		ids[i] = id
	}

	line, _ = reader.ReadString('\n')
	val, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)

	ans := findShortest(int32(graphNodes), graphFrom, graphTo, ids, val)
	fmt.Println(ans)
}
