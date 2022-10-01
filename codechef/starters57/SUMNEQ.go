package main

import (
	"fmt"
	"log"
)

func main() {
	var aliceInput int

	_, err := fmt.Scanf("%d", &aliceInput)

	if err != nil {
		log.Fatal(err)
	}

	if aliceInput <= 1 || aliceInput > 100 {
		fmt.Println(0)
	} else {
		fmt.Println(aliceInput - 1)
	}

}
