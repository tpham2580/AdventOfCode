package main

import (
	"aoc-go-2024/internal/logger"
	"bufio"
	"flag"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

type Service struct {
	log *logger.Logger
}

// Get the count of middle values from array of integers; Assume the arrays are always odd
func getMiddleCount(adj map[int][]int, arrInts [][]int) int {
	return 0
}

// Reads from file and return a adjacency map and an array of array of integers to process
func readFile(fileName string) (map[int][]int, [][]int) {
	slog.Info("Starting readFile")
	var adj = make(map[int][]int)
	var arrInt [][]int
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error(err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		regexNums := regexp.MustCompile(`[0-9]+`)

		line := scanner.Text()
		numbersStr := regexNums.FindAllString(line, -1)
		if len(numbersStr) == 2 {
			u, _ := strconv.Atoi(numbersStr[0])
			slog.Debug("Current key", "key", u)
			v, _ := strconv.Atoi(numbersStr[1])
			slog.Debug("Current value", "value", v)
			adj[u] = append(adj[u], v)
			slog.Debug("Current array for key", "arr", adj[u])
		} else {
			var numbers []int
			for i := range numbersStr {
				num, _ := strconv.Atoi(numbersStr[i])
				numbers = append(numbers, num)
			}

			if len(numbers) > 0 {
				slog.Debug("Page Updates", "update", numbers)
				arrInt = append(arrInt, numbers)
			}
		}
	}

	slog.Debug("readFile return", "adj", adj, "arrInt", arrInt)
	return adj, arrInt
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	// Flag
	file := flag.String("file", "day5input.txt", "Name of file that is being read")
	flag.Parse()

	adj, arrInt := readFile(*file)
	count := getMiddleCount(adj, arrInt)

	slog.Info("The answer to part one", "Middle Count", count)
}
