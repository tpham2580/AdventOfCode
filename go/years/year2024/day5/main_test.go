package main

import (
	"log/slog"
	"os"
	"testing"
)

func TestParts(t *testing.T) {
	testFile := "../../../../inputs/year2024/day5input_test.txt"
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	adj, arrInt := readFile(testFile)
	middleCount := getMiddleCount(adj, arrInt)

	t.Run("Part I", func(t *testing.T) {
		got := middleCount
		want := 143

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})

}
