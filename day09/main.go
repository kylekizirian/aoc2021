package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := readInput()
	part1(matrix)
}

func part1(matrix [][]int) {
	totalRisk := 0
	for row, rowSlice := range matrix {
		for col := range rowSlice {
			// check up, down, left, right
			if col > 0 && matrix[row][col-1] <= matrix[row][col] {
				continue
			}
			if row > 0 && matrix[row-1][col] <= matrix[row][col] {
				continue
			}
			if col < len(rowSlice)-1 && matrix[row][col+1] <= matrix[row][col] {
				continue
			}
			if row < len(matrix)-1 && matrix[row+1][col] <= matrix[row][col] {
				continue
			}
			totalRisk += matrix[row][col] + 1
		}
	}
	fmt.Println("part 1: ", totalRisk)
}

func readInput() [][]int {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var matrix [][]int
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "")
		nums := make([]int, len(vals))
		for i, val := range vals {
			nums[i] = strToInt(val)
		}
		matrix = append(matrix, nums)
	}

	return matrix
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
