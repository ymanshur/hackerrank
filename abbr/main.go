package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
 * Complete the 'abbreviation' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func abbreviation(a string, b string) string {
	var (
		na = len(a)
		nb = len(b)
	)

	if nb > na {
		return "NO"
	}

	valid := make([][]bool, na+1)
	valid[0] = make([]bool, nb+1)
	valid[0][0] = true
	for i := 1; i <= na; i++ {
		valid[i] = make([]bool, nb+1)

		if unicode.IsLower(rune(a[i-1])) {
			valid[i][0] = true
		}
	}

	for i := 1; i <= na; i++ {
		for j := 1; j <= nb; j++ {
			if a[i-1] == b[j-1] {
				valid[i][j] = valid[i-1][j-1]
				continue
			}

			if unicode.ToTitle(rune(a[i-1])) == rune(b[j-1]) {
				valid[i][j] = valid[i-1][j-1] || valid[i-1][j]
				continue
			}

			if unicode.IsUpper(rune(a[i-1])) {
				valid[i][j] = false
				continue
			}

			valid[i][j] = valid[i-1][j]
		}
	}

	//fmt.Printf("\t")
	//for j := 1; j <= nb; j++ {
	//    fmt.Printf("%c ", b[j-1])
	//}
	//fmt.Println()
	//
	//for i := 1; i <= na; i++ {
	//    fmt.Printf("%c\t", a[i-1])
	//    for j := 1; j <= nb; j++ {
	//        if valid[i][j] {
	//            fmt.Printf("%d ", 1)
	//        } else {
	//            fmt.Printf("%d ", 0)
	//        }
	//    }
	//    fmt.Println()
	//}

	if !valid[na][nb] {
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

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		a := readLine(reader)

		b := readLine(reader)

		result := abbreviation(a, b)

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
