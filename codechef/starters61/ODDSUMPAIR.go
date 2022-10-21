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

// https://www.codechef.com/START61D/problems/ODDSUMPAIR
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var a, b, c int
		scanf("%d %d %d\n", &a, &b, &c)

		if isPair(a) && isPair(b) && isPair(c) {
			printf("NO\n")
		} else if isOdd(a) && isOdd(b) && isOdd(c) {
			printf("NO\n")
		} else {
			printf("YES\n")
		}

	}

}

func isPair(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return !isPair(x)
}
