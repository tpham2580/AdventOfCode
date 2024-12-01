package main

import (
	"log"
	"os"
)

func main() {
	// TODO: Read line by line through day1input.txt, place first value into first list and second value into second list
	file, err := os.Open("day1input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var arr1 = [...]int{}
	var arr2 = [...]int{}

	// TODO: Sort the two lists

	// TODO: Go through both list at the same time, if the values are the same, add to running total

	// TODO: Output the running total

}
