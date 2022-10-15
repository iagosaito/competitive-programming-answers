package main

import (
	"bufio"
	"fmt"
	"math"
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

// https://www.codechef.com/OCT221D/problems/PODIUM
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var friends, slices int
		scanf("%d %d\n", &friends, &slices)

		totalOfSlices := friends * slices
		pizzas := math.Ceil(float64(totalOfSlices) / 4.0)

		printf("%d\n", int(pizzas))
	}

}
