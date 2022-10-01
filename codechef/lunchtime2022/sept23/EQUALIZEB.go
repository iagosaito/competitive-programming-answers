package main

import "fmt"

// Link: https://www.codechef.com/submit/EQUALIZEAB
// Não Concluída
func main() {
	var testCases int

	fmt.Scan(&testCases)

	for testCases > 0 {
		var a int
		var b int
		var x int
		fmt.Scanf("%d %d %d", &a, &b, &x)

		if isEqual(a, b, x) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		testCases--
	}
}

func isEqual(a, b, x int) bool {

	if a == b {
		return true
	}

	return a+x == b-x || a-x == b+x
}
