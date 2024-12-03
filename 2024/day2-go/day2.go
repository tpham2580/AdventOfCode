package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Both below need to be met to be considered safe:
	// The levels are either all increasing or all decreasing
	// Any two adjacent levels differ by at least one and at most three.

	file, err := os.Open("day2input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Go through the file line by line
	var safeCount int = 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		strArr := strings.Fields(line)
		intArr := make([]int, len(strArr))

		for i, s := range strArr {
			intArr[i], err = strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Issue with converting %d\n", s)
			}
		}

		var isIncreasing bool = false
		var isDecreasing bool = false
		var useDampener int = 0
		for i := 1; i < len(intArr); i++ {

			diff := intArr[i] - intArr[i-1]
			if diff >= 0 {
				if isDecreasing {
					useDampener++
					continue
				} else if diff < 1 || diff > 3 {
					useDampener++
				}
				isIncreasing = true
			} else {
				if isIncreasing {
					useDampener++
					continue
				} else if -diff < 1 || -diff > 3 {
					useDampener++
				}
				isDecreasing = true
			}

			if isIncreasing && isDecreasing {
				break
			}

			if useDampener > 1 {
				break
			}
		}

		if useDampener <= 1 {
			if isIncreasing && isDecreasing {
				continue
			} else if isIncreasing || isDecreasing {
				safeCount++
			}
		}
	}

	fmt.Printf("The number of safe levels is %d\n", safeCount)
	// Seperate all values and convert them all to int
	// Go through the whole list, and check if increasing each time
	// and also check if the difference between each step is between 1 to 3
	// If we get to the end, increment the safe counter

	// return the safe counter
}
