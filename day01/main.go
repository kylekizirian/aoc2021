package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	measurements := readInput()
	partA(measurements)
	partB(measurements)
}

func partA(measurements []int) {
	numIncreasing := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			numIncreasing++
		}
	}

	fmt.Println("part A: ", numIncreasing)
}

func partB(measurements []int) {
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

	fmt.Println("part b: ", numIncreasing)
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
