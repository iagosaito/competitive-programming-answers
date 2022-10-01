package main

import "fmt"

// Link: https://www.codechef.com/LTIME112D/problems/SLEEP
func main() {
	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var input int
		fmt.Scan(&input)

		if input < 7 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
