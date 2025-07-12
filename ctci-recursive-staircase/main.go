package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const M = 10000000007

/*
 * Complete the 'stepPerms' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func stepPerms(n int32) int32 {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if n == 3 {
		return 4
	}

	// var (
	//     curr,
	//     prev1,
	//     prev2,
	//     prev3,
	//     i int64
	// )

	// prev1, prev2, prev3 = 1, 2, 4
	// for i = 4; i <= int64(n); i++ {
	//     curr = prev1 + prev2 + prev3
	//     prev1, prev2, prev3 = prev2 % M, prev3 % M, curr % M
	// }

	// return int32(curr)

	mem := make([]int64, n+1)
	mem[0], mem[1], mem[2], mem[3] = 0, 1, 2, 4

	stepPermsMem(n, mem)

	return int32(mem[n])
}

func stepPermsMem(n int32, mem []int64) int64 {
	if mem[n] > 0 {
		return mem[n] % M
	}

	fmt.Println(n)

	mem[n] = stepPermsMem(n-int32(1), mem) + stepPermsMem(n-int32(2), mem) + stepPermsMem(n-int32(3), mem)

	return mem[n]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	s := int32(sTemp)

	for sItr := 0; sItr < int(s); sItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		res := stepPerms(n)

		fmt.Fprintf(writer, "%d\n", res)
	}

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
