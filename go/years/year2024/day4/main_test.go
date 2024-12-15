package main

import (
	"aoc-go-2024/internal/logger"
	"testing"
)

func TestParts(t *testing.T) {
	testFile := "../../../../inputs/year2024/day4input_test.txt"

	s := &Service{
		logger: logger.NewLogger(),
	}

	arrStrings := s.getStringArray(testFile)

	t.Run("Part I", func(t *testing.T) {
		got := s.getWordCount(arrStrings, "XMAS")
		want := 18

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})
	t.Run("Part II", func(t *testing.T) {
		got := s.getXmasCount(arrStrings)
		want := 9

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})

}
