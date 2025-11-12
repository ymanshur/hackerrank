package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Cookie struct {
	sweet int32
}

type PriorityCookies []*Cookie

func (p PriorityCookies) Len() int           { return len(p) }
func (p PriorityCookies) Less(i, j int) bool { return p[i].sweet < p[j].sweet }
func (p PriorityCookies) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PriorityCookies) Push(val any)      { *p = append(*p, val.(*Cookie)) }
func (p *PriorityCookies) Pop() any {
	n := len(*p)
	if n == 0 {
		return nil
	}

	tmp := *p
	val := tmp[n-1]
	tmp[n-1] = nil
	*p = tmp[:n-1]
	return val
}

/*
 * Complete the 'cookies' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 */

func cookies(k int32, A []int32) int32 {
	minOperation := int32(0)

	priorityA := make(PriorityCookies, len(A)) // min priority
	for i, sweet := range A {
		priorityA[i] = &Cookie{sweet: sweet}
	}

	heap.Init(&priorityA)

	var sweet1, sweet2, sweet3 int32
	for priorityA.Len() > 1 {
		cookie1 := heap.Pop(&priorityA).(*Cookie)
		sweet1 = cookie1.sweet
		if sweet1 >= k {
			break
		}

		cookie2 := heap.Pop(&priorityA).(*Cookie)
		sweet2 = cookie2.sweet

		sweet3 = sweet1 + 2*sweet2
		minOperation++

		heap.Push(&priorityA, &Cookie{sweet: sweet3})
	}

	if priorityA.Len() == 0 || priorityA[0].sweet < k {
		return -1
	}

	return minOperation
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var A []int32

	for i := 0; i < int(n); i++ {
		AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
		checkError(err)
		AItem := int32(AItemTemp)
		A = append(A, AItem)
	}

	result := cookies(k, A)

	fmt.Fprintf(writer, "%d\n", result)

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
