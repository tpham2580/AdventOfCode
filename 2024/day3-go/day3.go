package main

import (
	"aoc-go-2024/internal/logger"
	"os"
)

type Service struct {
	logger *logger.Logger
}

func (s *Service) readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		s.logger.Error(err.Error())
	}
	defer file.Close()

	var arrStrings []string

	return arrStrings
}

func main() {
	logger := logger.NewLogger()
	s := &Service{
		logger: logger,
	}

	fileInput := "day3input.txt"
	files := s.readFile(fileInput)
}
