package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func scanf(s string, v ...interface{}) {
	fmt.Fscanf(reader, s, v...)
}

func printf(s string, v ...interface{}) {
	fmt.Fprintf(writer, s, v...)
}

// https://www.codechef.com/START60D/problems/FLIPPAL
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var n int
		scanf("%d\n", &n)

		var s string
		scanf("%s\n", &s)

		czeros, cones := countsZeroAndOnes(s)

		if czeros%2 == 0 || cones%2 == 0 {
			printf("YES\n")
		} else {
			printf("NO\n")
		}
	}

}

func countsZeroAndOnes(b string) (int, int) {
	accZero := 0
	accOne := 0

	for _, v := range b {
		if x, _ := strconv.Atoi(string(v)); x == 0 {
			accZero++
		} else {
			accOne++
		}
	}

	return accZero, accOne
}
