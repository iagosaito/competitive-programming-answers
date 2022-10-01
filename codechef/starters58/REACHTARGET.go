package main

import "fmt"

// https://www.codechef.com/START58D/problems/REACHTARGET
func main() {
	var testCases int

	fmt.Scan(&testCases)

	for testCases > 0 {
		var x int
		var y int
		fmt.Scanf("%d %d", &x, &y)

		fmt.Println(x - y)

		testCases--
	}
}
