package main

import (
	"fmt"
)

// Link: https://www.codechef.com/LTIME112D/problems/FLIPCARDS
// CORRECT
func main() {
	var t int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var l int
		var k int

		fmt.Scanf("%d %d", &l, &k)

		if l == k || l%k == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}

// Link: https://www.codechef.com/LTIME112D/problems/STICKBREAK
// Incorrect
// func main() {
// 	var t int

// 	fmt.Scan(&t)

// 	for i := 0; i < t; i++ {
// 		var l int
// 		var k int

// 		fmt.Scanf("%d %d", &l, &k)

// 		if l == k || l%k == 0 {
// 			fmt.Println(0)
// 			continue
// 		}

// 		v := int(math.Round(float64(l) / float64(k)))

// 		arrayOfSticks := make([]int, k-1)

// 		acc := 0

// 		for i := 0; i < k-1; i++ {
// 			arrayOfSticks[i] = v
// 			acc = v + acc
// 		}

// 		diff := acc - l

// 		fmt.Println("diff:", diff, acc)

// 		if diff > 0 {
// 			for i := 0; i < len(arrayOfSticks); i++ {
// 				if diff <= 0 {
// 					break
// 				}

// 				arrayOfSticks[i] = arrayOfSticks[i] - 1
// 				diff--
// 			}
// 		} else {
// 			for i := 0; i < len(arrayOfSticks); i++ {
// 				if diff >= 0 {
// 					break
// 				}

// 				arrayOfSticks[i] = arrayOfSticks[i] + 1
// 				diff++
// 			}
// 		}

// 		result := 0

// 		for i := 0; i < len(arrayOfSticks)-1; i++ {
// 			fmt.Println(arrayOfSticks[i])
// 			r := arrayOfSticks[i] - arrayOfSticks[i+1]

// 			if r < 0 {
// 				result += r * -1
// 			} else {
// 				result += r
// 			}
// 		}

// 		fmt.Println(result)
// 	}
// }
