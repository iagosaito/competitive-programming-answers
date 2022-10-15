package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var alicePicks string
var bobPicks string

func scanf(s string, v ...interface{}) {
	fmt.Fscanf(reader, s, v...)
}

func printf(s string, v ...interface{}) {
	fmt.Fprintf(writer, s, v...)
}

// https://www.codechef.com/START59D/problems/SUSSTR
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {

		var n, seconds int
		scanf("%d %d\n", &n, &seconds)

		var input string
		scanf("%s\n", &input)

		r := ""
		for seconds > 0 {
			seconds--

			for i := 0; i < len(input); i++ {
				if string(v) == "1" {
					r := r + 0
				}
			}
		}

		fmt.Println(input)
	}
}
