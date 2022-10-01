package main

import (
	"fmt"
	"math"
)

// https://www.codechef.com/START58D/problems/RANKLISTPAGE
func main() {
	var testCases int

	fmt.Scan(&testCases)

	for testCases > 0 {
		var rank int
		fmt.Scanf("%d", &rank)

		fmt.Println(math.Ceil(float64(rank) / 25.0))

		testCases--
	}
}
