package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var vowels = map[string]bool{
	"a": true, "e": true, "i": true, "o": true, "u": true,
}

func scanf(s string, v ...interface{}) {
	fmt.Fscanf(reader, s, v...)
}

func printf(s string, v ...interface{}) {
	fmt.Fprintf(writer, s, v...)
}

// https://www.codechef.com/START59D/problems/HAPPYSTR
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

outerfor:
	for i := 0; i < testCases; i++ {
		var x string
		scanf("%s\n", &x)

		lowerx := strings.ToLower(x)

		i := 0
		for i < len(lowerx)-3 {
			i++
			if vowels[string(lowerx[i])] == false {
				continue
			}

			subx := lowerx[i : i+3]
			if isVowel(subx) {
				printf("Happy\n")
				continue outerfor
			}
		}

		printf("Sad\n")
	}
}

func isVowel(x string) bool {

	for _, v := range x {
		_, r := vowels[string(v)]
		if r == false {
			return false
		}
	}

	return true
}
