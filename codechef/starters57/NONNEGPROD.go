package main

import (
	"fmt"
)

func main() {

	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var arraySize int
		fmt.Scan(&arraySize)

		array := read(arraySize)

		negativeNumbers := 0

		for i := 0; i < arraySize; i++ {
			if array[i] == 0 {
				negativeNumbers = 0
				break
			}

			if array[i] < 0 {
				negativeNumbers += 1
			}
		}

		if negativeNumbers%2 == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}

	}

}

func read(n int) []int {
	in := make([]int, n)
	for i := range in {
		fmt.Scan(&in[i])
	}
	return in
}
