package main

import (
	"aoc-go-2024/internal/logger"
	"bufio"
	"flag"
	"os"
)

type Service struct {
	log *logger.Logger
}

// Reads from file and return a adjacency map and an array of array of integers to process
func (s *Service) readFile(fileName string) (map[int][]int, [][]int) {
	var adj = make(map[int][]int)
	var arrInt [][]int
	file, err := os.Open(fileName)
	if err != nil {
		s.log.Error(err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

	}

	return adj, arrInt
}

func main() {
	s := &Service{
		log: logger.NewLogger(),
	}

	// Flag
	file := flag.String("file", "day5input.txt", "Name of file that is being read")
	flag.Parse()

	adj, arrInt := s.readFile(*file)
}
