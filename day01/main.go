package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	measurements := readInput()
	part1(measurements)
	part2(measurements)
}

func part1(measurements []int) {
	numIncreasing := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			numIncreasing++
		}
	}

	fmt.Println("part 1: ", numIncreasing)
}

func part2(measurements []int) {
	numIncreasing := 0
	firstThree := measurements[0] + measurements[1] + measurements[2]
	secondThree := measurements[1] + measurements[2] + measurements[3]
	if secondThree > firstThree {
		numIncreasing++
	}

	for i := 4; i < len(measurements); i++ {
		firstThree = firstThree + measurements[i-1] - measurements[i-4]
		secondThree = secondThree + measurements[i] - measurements[i-3]
		if secondThree > firstThree {
			numIncreasing++
		}
	}

	fmt.Println("part 2: ", numIncreasing)
}

func readInput() []int {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var measurements []int
	for scanner.Scan() {
		measure, err := strconv.Atoi(scanner.Text())
		checkErr(err)
		measurements = append(measurements, measure)
	}

	return measurements
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
