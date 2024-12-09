package main

import "aoc-go-2024/internal/logger"

type Service struct {
	logger *logger.Logger
}

func main() {
	logger := logger.NewLogger()
	s := &Service{
		logger: logger,
	}

	fileInput := "day3input.txt"
}
