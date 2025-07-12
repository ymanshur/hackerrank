package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack []rune

func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	idx := len(*s) - 1
	val := (*s)[idx]
	(*s)[idx] = 0
	*s = (*s)[:idx]
	return val, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

/*
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isBalanced(s string) string {
	var stack Stack
	for _, val := range s {
		switch val {
		case '}':
			b, _ := stack.Pop()
			if b != '{' {
				return "NO"
			}
		case ')':
			b, _ := stack.Pop()
			if b != '(' {
				return "NO"
			}
		case ']':
			b, _ := stack.Pop()
			if b != '[' {
				return "NO"
			}
		default:
			stack.Push(val)
		}
	}

	if !stack.IsEmpty() {
		return "NO"
	}

	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Fprintf(writer, "%s\n", result)
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
