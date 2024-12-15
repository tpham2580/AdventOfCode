package main

import (
	"aoc-go-2024/internal/logger"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Service struct contains logger
type Service struct {
	logger *logger.Logger
}

func abs(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

func (s *Service) readReports(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		s.logger.Error(err.Error())
	}
	defer file.Close()

	var reports = [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var report []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				s.logger.Error(err.Error())
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	return reports
}

func (s *Service) isSafe(report []int) bool {
	var increasing bool = true
	var decreasing bool = true

	for i := 0; i < len(report)-1; i++ {
		var diff int = report[i+1] - report[i]

		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if diff < 0 {
			increasing = false
		}
		if diff > 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func (s *Service) countSafeReportsWithDampener(reports [][]int) int {
	var safeCount int = 0

	for _, report := range reports {
		if s.isSafe(report) {
			safeCount++
		} else {
			for index := 0; index < len(report); index++ {
				leftPart := append([]int{}, report[:index]...)
				rightPart := append([]int{}, report[index+1:]...)
				modifiedReport := append(leftPart, rightPart...)
				if s.isSafe(modifiedReport) {
					safeCount++
					break
				}
			}
		}
	}

	return safeCount
}

func (s *Service) countSafeReports(reports [][]int) int {
	var safeCount int = 0

	for _, report := range reports {
		isSafeReport := s.isSafe(report)
		if isSafeReport {
			safeCount++
		}
	}

	return safeCount
}

func main() {
	logger := logger.NewLogger()
	s := &Service{
		logger: logger,
	}

	// Flags
	useDampener := flag.Bool("dampener", false, "Use dampener to calculate safe reports")
	flag.Parse()

	fileInput := "day2input.txt"
	reports := s.readReports(fileInput)

	var safeCount int
	if *useDampener {
		s.logger.Info("Calculating safe reports with dampener...")
		safeCount = s.countSafeReportsWithDampener(reports)
	} else {
		s.logger.Info("Calculating safe reports without dampener...")
		safeCount = s.countSafeReports(reports)
	}

	s.logger.Info(fmt.Sprintf("Safe count: %v", safeCount))
}
