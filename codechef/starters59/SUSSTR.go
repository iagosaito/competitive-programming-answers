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

func isEven(n int) bool {
	return n%2 == 0
}

// https://www.codechef.com/START59D/problems/SUSSTR
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		alicePicks = ""
		bobPicks = ""
		var n int
		scanf("%d\n", &n)

		var binaryString string
		scanf("%s\n", &binaryString)

		if isEven(n) {
			alicePicks = binaryString[:(len(binaryString) / 2)]
			bobPicks = binaryString[(len(binaryString) / 2):]
		} else {
			alicePicks = binaryString[:(len(binaryString)/2)+1]
			bobPicks = binaryString[(len(binaryString)/2)+1:]
		}

		r := ""
		x := 0
		y := len(bobPicks) - 1
		for x < len(alicePicks) {
			if string(alicePicks[x]) == "1" {
				r = r + string(alicePicks[x])
			} else {
				r = string(alicePicks[x]) + r
			}

			x++

			for y >= 0 {
				if string(bobPicks[y]) == "1" {
					r = string(bobPicks[y]) + r
				} else {
					r = r + string(bobPicks[y])
				}
				break
			}

			y--
		}

		printf("%s\n", r)
	}
}
