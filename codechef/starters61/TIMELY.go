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

// https://www.codechef.com/START61D/problems/TIMELY
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var x int
		scanf("%d\n", &x)

		if x < 30 {
			printf("NO\n")
		} else {
			printf("YES\n")
		}

	}

}
