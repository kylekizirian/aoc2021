package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	timers := readInput()

	due := make([]int, 300)

	for _, timer := range timers {
		due[timer]++
	}

	numFish := len(timers)
	for i := 0; i < 256; i++ {
		if i == 80 {
			fmt.Println("part 1: ", numFish)
		}

		numFish += due[i]
		due[i+7] += due[i]
		due[i+9] += due[i]
	}

	fmt.Println("part 2: ", numFish)
}

func readInput() []int {
	file, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fields := strings.Split(scanner.Text(), ",")
	timers := make([]int, len(fields))

	for i, field := range fields {
		timers[i] = strToInt(field)
	}

	return timers
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
