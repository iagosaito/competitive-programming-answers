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

// https://www.codechef.com/START59D/problems/SPECIALITY
func main() {

	defer writer.Flush()

	var testCases int

	scanf("%d\n", &testCases)

	for i := 0; i < testCases; i++ {
		var setter, tester, editorialist int
		scanf("%d %d %d\n", &setter, &tester, &editorialist)

		if setter > tester && setter > editorialist {
			printf("Setter\n")
		} else if tester > setter && tester > editorialist {
			printf("Tester\n")
		} else {
			printf("Editorialist\n")
		}
	}

}
