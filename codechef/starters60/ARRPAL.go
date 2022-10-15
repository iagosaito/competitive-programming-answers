package main

import (
	"bufio"
	"fmt"
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

// https://www.codechef.com/START60D/problems/ARRPAL
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var n int
		scanf("%d\n", &n)

		l := make([]int, n)
		_, posHigh, occ := scanList(n, l)

		if occ == n || isPalindrome(l) {
			printf("%d\n", 0)
		} else if posHigh < n {
			printf("%d\n", -1)
		} else {
			printf("%d\n", l[n-1]-l[0])
		}

	}

}

func isPalindrome(input []int) bool {
	firstHalf := make([]int, len(input))
	secondHalf := make([]int, len(input))

	half := len(input) / 2
	if len(input)%2 == 0 {
		firstHalf = input[0:half]
		secondHalf = input[half:]
	} else {
		firstHalf = input[0 : half+1]
		secondHalf = input[half:]
	}

	reverseSecondHalf := reverseInts(secondHalf)

	for i, v := range firstHalf {
		if v != reverseSecondHalf[i] {
			return false
		}
	}

	return true
}

func reverseInts(input []int) []int {

	var newSlice []int
	for i := 0; i < len(input); i++ {
		newSlice = append(newSlice, input[i])
	}

	if len(newSlice) == 0 {
		return newSlice
	}

	i := 0
	j := len(newSlice) - 1

	for i < j {
		newSlice[i], newSlice[j] = newSlice[j], newSlice[i]
		i++
		j--
	}

	return newSlice
}

func scanList(n int, m []int) (int, int, int) {

	high := 0
	indiceHigh := 0
	occurences := 0

	for i := 0; i < n; i++ {
		var x int
		scanf("%d ", &x)

		m[i] = x

		if x > high {
			high = x
			indiceHigh = i + 1
			occurences = 1
		} else if x == high {
			occurences++
		}
	}

	return high, indiceHigh, occurences
}
