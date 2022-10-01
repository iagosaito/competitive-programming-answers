package main

import (
	"fmt"
	"log"
)

func main() {
	var aliceMark int
	var bobMark int

	_, err := fmt.Scanf("%d %d", &aliceMark, &bobMark)

	if err != nil {
		log.Fatal(err)
	}

	if bobMark*2 <= aliceMark {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
