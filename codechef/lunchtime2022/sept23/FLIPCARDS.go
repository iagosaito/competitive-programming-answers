package main

import "fmt"

// Link: https://www.codechef.com/LTIME112D/problems/FLIPCARDS
func main() {
	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var totalOfCards int
		var facedUpCards int
		fmt.Scanf("%d %d", &totalOfCards, &facedUpCards)

		if totalOfCards == facedUpCards {
			fmt.Println(0)
		} else if totalOfCards-facedUpCards >= facedUpCards {
			fmt.Println(facedUpCards)
		} else {
			fmt.Println(totalOfCards - facedUpCards)
		}

	}
}
