package main

import (
	"aoc-go-2024/internal/logger"
	"bufio"
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

// Uses dfs to search for word in board. Returns count of word.
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

func main() {
	s := &Service{
		logger: logger.NewLogger(),
	}

	fileName := "day4input.txt"
	arrStrings := s.getStringArray(fileName)
	wordCount := s.getWordCount(arrStrings, "XMAS")
	s.logger.Info(strconv.Itoa(wordCount))
}
