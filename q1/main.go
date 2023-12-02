package main

import (
	"advent_of_code_q1/parser"
	"fmt"
)

func main() {
	// Injest the secret file
	n, err := parser.Parse("../secret/q1_1", parser.ParseSum)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The sum for part 1 is: %d\n", n)

	n, err = parser.Parse("../secret/q1_1", parser.ParseSumWithStrings)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The sum for part 2 is: %d\n", n)
}
