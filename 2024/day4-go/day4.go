package main

import (
	"aoc-go-2024/internal/logger"
	"bufio"
	"flag"
	"os"
	"strconv"
)

type Service struct {
	logger *logger.Logger
}

type Direction struct {
	X, Y int
}

type Coordinate struct {
	row, col int
}

func (s *Service) getStringArray(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		s.logger.Error(err.Error())
	}

	var arrStrings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arrStrings = append(arrStrings, line)
	}

	return arrStrings
}

// Uses dfs to search in 8 directions for word in board. Returns count of word.
func (s *Service) getWordCount(board []string, word string) int {
	var ROWS int = len(board)
	var COLS int = len(board[0])
	var count int = 0
	var directions = []Direction{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

	// Create 2d Mutable Board to mark coordinates through DFS
	mutBoard := make([][]rune, ROWS)
	for i := range board {
		mutBoard[i] = []rune(board[i])
	}

	s.logger.Info("Searching for " + word)

	// Checks to see if the points are valid and if the next character is the next word[index]
	isValid := func(row, col, index int) bool {
		return (0 <= row && row < ROWS) && (0 <= col && col < COLS) && mutBoard[row][col] == rune(word[index])
	}

	// DFS function that recursively looks a word
	var findWord func(row, col, index int, direction Direction)
	findWord = func(row, col, index int, direction Direction) {
		if index == len(word) {
			count++
			return
		}
		if isValid(row, col, index) {
			temp := mutBoard[row][col]
			mutBoard[row][col] = '#'
			findWord(row+direction.X, col+direction.Y, index+1, direction)
			mutBoard[row][col] = temp
		}
	}

	// go through entire board and look search for word at each point
	for row := range ROWS {
		for col := range COLS {
			if mutBoard[row][col] == rune(word[0]) {
				for _, direction := range directions {
					findWord(row, col, 0, direction)
				}
			}
		}
	}
	return count
}

func (s *Service) getXmasCount(board []string) int {
	var ROWS int = len(board)
	var COLS int = len(board[0])
	var count int = 0
	const A = 'A'

	s.logger.Info("Searching for X-MAS")

	var isValidXmas func(row, col int) bool
	isValidXmas = func(row, col int) bool {
		topLeft, topRight, botLeft, botRight := board[row-1][col-1], board[row+1][col-1], board[row-1][col+1], board[row+1][col+1]

		if (topLeft == 'M' && botRight == 'S') || (topLeft == 'S' && botRight == 'M') {
			if (topRight == 'M' && botLeft == 'S') || (topRight == 'S' && botLeft == 'M') {
				return true
			}
		}

		return false
	}

	// go through entire board and look search for word at each point
	for row := 1; row < ROWS-1; row++ {
		for col := 1; col < COLS-1; col++ {
			if board[row][col] == A {
				if isValidXmas(row, col) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	s := &Service{
		logger: logger.NewLogger(),
	}

	// Flags part2
	xmas := flag.Bool("part2", false, "Checks for number of X-MAS instead because the elf couldn't explain the first time")
	word := flag.String("word", "XMAS", "Pass in word that will be searched from input text file")
	file := flag.String("file", "day4input.txt", "Pass in file name to read from file")
	flag.Parse()

	arrStrings := s.getStringArray(*file)

	var count int
	if *xmas {
		count = s.getXmasCount(arrStrings)
	} else {
		count = s.getWordCount(arrStrings, *word)
	}
	s.logger.Info(strconv.Itoa(count))
}
