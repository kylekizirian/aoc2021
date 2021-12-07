package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	positions := readInput()
	part1(positions)
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
	fmt.Println(dist)
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
