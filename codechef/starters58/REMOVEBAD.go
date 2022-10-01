package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func scanf(s string, v ...interface{}) {
	fmt.Fscanf(reader, s, v...)
}

func printf(s string, v ...interface{}) {
	fmt.Fprintf(writer, s, v...)
}

// https://www.codechef.com/START58D/problems/REMOVEBAD
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var arraySize int
		scanf("%d\n", &arraySize)

		m := make(map[int]int, arraySize)
		high := readList(arraySize, m)

		printf("%d\n", arraySize-high)
	}

}

func readList(n int, m map[int]int) (high int) {

	high = 0

	for i := 0; i < n; i++ {
		var n int
		scanf("%d ", &n)

		m[n] = m[n] + 1

		if m[n] > high {
			high = m[n]
		}
	}

	return high
}
