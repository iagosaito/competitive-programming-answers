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
		var currentFloorChef, currentFloorChefina, speedChef, speedChefina int
		scanf("%d %d %d %d\n", &currentFloorChef, &currentFloorChefina, &speedChef, &speedChefina)

		chefMin := minutesToGroundFloor(currentFloorChef, speedChef)
		chefinaMin := minutesToGroundFloor(currentFloorChefina, speedChefina)

		if chefMin == chefinaMin {
			printf("Both\n")
		} else if chefMin > chefinaMin {
			printf("Chefina\n")
		} else {
			printf("Chef\n")
		}
	}

}

func minutesToGroundFloor(currentFloor, speed int) int {
	temp := float64(currentFloor) / float64(speed)
	rounded := math.Round(temp*100) / 100

	return int(rounded * 100)
}
