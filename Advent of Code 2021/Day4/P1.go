package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var called_nums []string

	var boards [100][5][]string
	var hitBoards [][][]bool

	var sum = 0

	var winning_number int

	bingo_boardsF, _ := os.Open("bingo.txt")

	scanner := bufio.NewScanner(bingo_boardsF)

	scanner.Scan()
	called_nums = strings.Split(scanner.Text(), ",")

	scanner.Scan()
	lineInd := 0
	boardInd := 0
	for scanner.Scan() {
		var line []string
		if scanner.Text() != "" {
			lineInd++
			line = strings.Split(scanner.Text(), " ")
			for x := 0; x < 5; x++ {
				boards[boardInd][x] = line
			}
			if lineInd%5 == 0 {
				boardInd++
			}
		}
	}

out:
	for true == true {
		for i := 0; i < len(called_nums); i++ {
			for z := 0; z < len(boards); z++ {
				for x := 0; x < 5; x++ {
					for y := 0; y < 5; y++ {
						if boards[z][x][y] == called_nums[i] {
							hitBoards[z][x][y] = true
						}
					}
				}
			}
			for z := 0; z < len(hitBoards); z++ {
				isComplete, isRow, Ind := CheckBoard(hitBoards[z])
				winning_number = i
				if isComplete {
					winningBoard := boards[z]
					if isRow {
						for c := 0; c < 5; c++ {
							winningBoard[c][Ind] = ""
						}
					} else {
						for c := 0; c < 5; c++ {
							winningBoard[Ind][c] = ""
						}
					}

					for c := 0; c < 5; c++ {
						for d := 0; d < 5; d++ {
							iWinningBoard, _ := strconv.Atoi(winningBoard[c][d])
							sum += iWinningBoard
						}
					}

					break out

				}
			}
		}

	}

	println(sum * winning_number)
}

func CheckBoard(board [][]bool) (bool, bool, int) {
	var isComplete bool = false
	var Ind int
	var isRow bool = false

	for x := 0; x < 5; x++ {
		for y := 1; y < 5; y++ {
			if board[x][y] != board[x][0] {
				break
			} else {
				Ind = x
				isComplete = true
			}
		}
	}

	if !isComplete {
		for y := 0; y < 5; y++ {
			for x := 1; x < 5; x++ {
				if board[x][y] != board[0][y] {
					break
				} else {
					Ind = y
					isComplete = true
					isRow = true
				}
			}
		}

	}

	return isComplete, isRow, Ind
}
