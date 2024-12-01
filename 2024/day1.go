package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// TODO: Read line by line through day1input.txt, place first value into first list and second value into second list
	file, err := os.Open("day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr1 = []int{}
	var arr2 = []int{}

	var lineNumber int = 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber += 1

		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			log.Fatalf("Expected 2 numbers per line, but got %d", len(numbers))
		}

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}

		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}

		arr1 = append(arr1, num1)
		arr2 = append(arr2, num2)
	}

	// TODO: Sort the two lists
	sort.Ints(arr1)
	sort.Ints(arr2)

	// TODO: Go through both list at the same time, if the values are the same, add to running total
	var total int = 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] < arr2[i] {
			total += arr2[i] - arr1[i]
		} else {
			total += arr1[i] - arr2[i]
		}
	}

	// TODO: Output the running total
	fmt.Printf("The total distance between left and right lists is: %d\n", total)
}
