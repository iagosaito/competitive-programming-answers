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

// https://www.codechef.com/START60D/problems/TAXES
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var x, y int
		scanf("%d %d\n", &x, &y)

		x = 100 * x
		y = 100 * y

		x = x / 2

		if x <= y {
			printf("YES\n")
		} else {
			printf("NO\n")
		}
	}

}
