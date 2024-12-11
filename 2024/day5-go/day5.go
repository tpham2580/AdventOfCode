package main

import (
	"aoc-go-2024/internal/logger"
	"flag"
)

type Service struct {
	log *logger.Logger
}

func main() {
	s := &Service{
		log: logger.NewLogger(),
	}

	// Flag
	file := flag.String("file", "day5input.txt", "Name of file that is being read")
	flag.Parse()
}
