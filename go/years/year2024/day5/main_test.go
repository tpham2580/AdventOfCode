package main

import (
	"aoc-go-2024/internal/logger"
	"testing"
)

func TestParts(t *testing.T) {
	testFile := "../../../../inputs/year2024/day5input_test.txt"

	s := &Service{
		log: logger.NewLogger(),
	}

	aadj, arrInt := s.readFile(testFile)

	t.Run("Part I", func(t *testing.T) {
		got := 142
		want := 143

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})

}
