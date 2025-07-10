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
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

// Naive approach ~> O(n * r) time, O(1) space
// a = {1 2 3 4 5}, n, r := len(a), d % n
//      i (idx)
// a = {2 3 4 5 5} -> keep a[0] in tmp; change a[i] by a[i+1]; change the a[n-1] by tmp
// a = {3 4 5 1 2}
// a = {4 5 1 2 3} -> done with r iteration

// Better approach ~> O(n) time, O(n) space
// 1. let a2 with len(a) capacity
// 2. copy the a to a2,   from [r] to [n-1]
// 3. append the a to a2, from [0] to [r-1]
//
// a = {1 2 3 4 5} -> a2 = {4 5}
// a = {1 2 3}     -> a2 = {4 5} + {1 2 3}

// Expected approach ~> O(n) time, O(1) space
// a = {1 2 3 4 5} -> reverse from a[0] to a[r-1]
// a = {3 2 1 5 4} -> reverse from a[r] to a[n-1]
// a = {4 5 1 2 3} -> reverse from a[0] to a[n-1]

func rotLeft(a []int32, d int32) []int32 {
	var (
		n = int32(len(a))
		r = d % n
	)

	// for r > 0 {
	//     tmp := a[0]
	//     var i int32
	//     for i = 1; i < n; i++ {
	//         a[i-1] = a[i]
	//     }
	//     a[n-1] = tmp
	//     r--
	// }

	// a2 := make([]int32, n)
	// a2 = a[r:n]
	// a2 = append(a2, a[:r]...)

	for i, j := int32(0), r-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for i, j := r, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for i, j := int32(0), n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return a
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

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := rotLeft(a, d)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
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
