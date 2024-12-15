package main

import (
	"aoc-go-2024/internal/logger"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Service struct {
	logger *logger.Logger
}

func (s *Service) sumArray(nums []int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Reads file, gets all matches for mul(num1, num2), multiplies num1 and num2, and returns array of multiplied ints
func (s *Service) readFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		s.logger.Error(err.Error())
	}
	defer file.Close()

	var arr []int
	var do bool = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// regex to check for mul(num,num), do(), and don't()
		regexReal := regexp.MustCompile(`(mul\([0-9]+,[0-9]+\)|do\(\)|don\'t\(\))`)
		regexNum := regexp.MustCompile(`[0-9]+`)

		line := scanner.Text()
		calls := regexReal.FindAllString(line, -1) // Get all strings that are real

		for _, match := range calls {
			nums := regexNum.FindAllString(match, -1)

			if match == "do()" {
				do = true
			} else if match == "don't()" {
				do = false
			} else if len(nums) == 2 && do {
				nums1, _ := strconv.Atoi(nums[0])
				nums2, _ := strconv.Atoi(nums[1])
				arr = append(arr, nums1*nums2) // multiply both nums and add to array
			}
		}
	}
	return arr
}

func main() {
	logger := logger.NewLogger()
	s := &Service{
		logger: logger,
	}

	fileInput := "day3input.txt"
	arrMul := s.readFile(fileInput)
	var total int = s.sumArray(arrMul)
	s.logger.Info(strconv.Itoa(total))
}
