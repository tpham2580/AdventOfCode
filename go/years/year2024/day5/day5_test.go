package main

import (
	"aoc-go-2024/internal/logger"
	"testing"
)

func TestMain(m *testing.M) {
	testFile := "day5input_test.txt"

	s := &Service{
		logger.NewLogger(),
	}

	testLine, err := s.readFile(testFile)
	if err != nil {
		s.log.Error(err)
	}

	m.Run()
}

func TestPart1(t *testing.T) {

}
