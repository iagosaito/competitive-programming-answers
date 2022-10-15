package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func scanf(s string, v ...interface{}) {
	fmt.Fscanf(reader, s, v...)
}

func printf(s string, v ...interface{}) {
	fmt.Fprintf(writer, s, v...)
}

// https://www.codechef.com/OCT221D/problems/PODIUM
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var binLength int
		scanf("%d\n", &binLength)

		var s string
		scanf("%s\n", &s)

		i, _ := strconv.ParseInt(s, 2, 64)

		fmt.Println("Value of i:", i)

		y := flip(s)

		fmt.Println("Y flipado:", y)

		r, _ := strconv.ParseInt(y, 2, 64)
		fmt.Println("Value of r:", r)

		if r == 0 {
			fmt.Println("É zero")
			printf("%d\n", i)
		} else {
			x := i / r
			printf("%d\n", x)
		}

	}

}

func flip(s string) string {
	replacer := strings.NewReplacer("0", "1", "1", "0")
	return replacer.Replace(s)
}
