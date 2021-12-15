package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := readInput()
	part1(grid)
}

func part1(grid [][]int) {
	minRisk := make([][]int, len(grid))
	for i := range minRisk {
		minRisk[i] = make([]int, len(grid[i]))
	}

	for i := range minRisk {
		for j := range minRisk[i] {
			if i == 0 && j == 0 {
				continue
			}
			minimum := math.MaxInt
			if i > 0 {
				minimum = minRisk[i-1][j] + grid[i][j]
			}
			if j > 0 && minRisk[i][j-1]+grid[i][j] < minimum {
				minimum = minRisk[i][j-1] + grid[i][j]
			}
			minRisk[i][j] = minimum
		}
	}

	lastRow := len(minRisk) - 1
	lastCol := len(minRisk[0]) - 1
	fmt.Println("part 1: ", minRisk[lastRow][lastCol])
}

func readInput() [][]int {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var output [][]int

	for scanner.Scan() {
		var line []int
		for _, val := range strings.Split(scanner.Text(), "") {
			line = append(line, strToInt(val))
		}
		output = append(output, line)
	}

	return output
}

func strToInt(val string) int {
	num, err := strconv.Atoi(val)
	checkErr(err)
	return num
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
