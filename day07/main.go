package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	positions := readInput()
	part1(positions)
	part2(positions)
}

func part1(positions []int) {
	sort.Slice(positions, func(i, j int) bool {
		return positions[i] < positions[j]
	})

	med := positions[len(positions)/2]

	dist := 0
	for _, pos := range positions {
		dist += absInt(pos - med)
	}

	fmt.Println("part 1: ", dist)
}

func part2(positions []int) {
	minStart, maxStart := minAndMax(positions)
	minDistCost := math.MaxInt

	for i := minStart; i <= maxStart; i++ {
		distCost := 0

		for _, pos := range positions {
			distCost += fuelCost(i, pos)
		}

		if distCost < minDistCost {
			minDistCost = distCost
		}
	}

	fmt.Println("part 2: ", minDistCost)
}

func minAndMax(nums []int) (int, int) {
	if len(nums) == 0 {
		panic("empty slice")
	}

	min := math.MaxInt
	max := math.MinInt

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func fuelCost(a, b int) int {
	n := absInt(a - b)
	return n * (n + 1) / 2
}

func readInput() []int {
	file, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fields := strings.Split(scanner.Text(), ",")
	positions := make([]int, len(fields))

	for i, field := range fields {
		positions[i] = strToInt(field)
	}

	return positions
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func strToInt(num string) int {
	val, err := strconv.Atoi(num)
	checkErr(err)
	return val
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
