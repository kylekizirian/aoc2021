package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	draws, boards := readInput()
	part1(draws, boards)
	part2(draws, boards)
}

func part1(draws []int, boards []*board) {
	for _, draw := range draws {
		for _, board := range boards {
			board.markNum(draw)
			if board.hasBingo() {
				fmt.Println("part 1: ", board.unmarkedSum()*draw)
				return
			}
		}
	}
}

func part2(draws []int, boards []*board) {
	for _, draw := range draws {
		var notWon []*board
		for _, board := range boards {
			board.markNum(draw)
			if !board.hasBingo() {
				notWon = append(notWon, board)
			} else if len(boards) == 1 {
				// last board has won
				fmt.Println("part 2: ", board.unmarkedSum()*draw)
				return
			}
		}
		boards = notWon
	}
}

type boardNum struct {
	num    int
	marked bool
}

type board struct {
	nums [5][5]boardNum
}

func newBoard(lines []string) *board {
	// convert each line into len 5 array
	var b board
	for i := 0; i < 5; i++ {
		for j, numStr := range strings.Fields(lines[i]) {
			num, err := strconv.Atoi(numStr)
			checkErr(err)
			b.nums[i][j] = boardNum{num, false}
		}
	}
	return &b
}

func (b *board) String() string {
	var res string
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.nums[i][j].marked {
				res += fmt.Sprintf("*%2d", b.nums[i][j].num)
			} else {
				res += fmt.Sprintf("%3d", b.nums[i][j].num)
			}
		}
		res += "\n"
	}
	return res
}

func (b *board) markNum(num int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.nums[i][j].num == num {
				b.nums[i][j].marked = true
			}
		}
	}
}

func (b *board) hasBingo() bool {
	for i := 0; i < 5; i++ {
		bingo := true
		for j := 0; j < 5; j++ {
			if !b.nums[i][j].marked {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		bingo := true
		for j := 0; j < 5; j++ {
			if !b.nums[j][i].marked {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}

	return false
}

func (b *board) unmarkedSum() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.nums[i][j].marked {
				sum += b.nums[i][j].num
			}
		}
	}
	return sum
}

func readInput() ([]int, []*board) {
	// line will be a list of ints, separated by commas
	file, err := os.Open("input.txt")
	checkErr(err)

	var drawnNums []int
	var boards []*board

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	for _, numStr := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(numStr)
		checkErr(err)
		drawnNums = append(drawnNums, num)
	}

	var lines []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			lines = []string{}
		} else {
			lines = append(lines, scanner.Text())
			if len(lines) == 5 {
				board := newBoard(lines)
				boards = append(boards, board)
			}
		}
	}

	return drawnNums, boards
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
