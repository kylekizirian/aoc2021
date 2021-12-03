package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readInput()
	part1(lines)
}

func part1(binNums []string) {
	// increment when we see a 1, decrement on 0
	counters := make([]int, len(binNums[0]))
	for pos := range binNums[0] {
		for _, binNum := range binNums {
			switch string(binNum[pos]) {
			case "0":
				counters[pos] -= 1
			case "1":
				counters[pos] += 1
			default:
				panic("unexpected digit in binarys tring")
			}
		}
	}

	gammaDigits := ""
	epsilonDigits := ""

	for _, counter := range counters {
		if counter < 0 {
			gammaDigits += "0"
			epsilonDigits += "1"
		} else if counter > 0 {
			gammaDigits += "1"
			epsilonDigits += "0"
		} else {
			panic("same number of 0 and 1s")
		}
	}

	gamma, err := strconv.ParseInt(gammaDigits, 2, 64)
	checkErr(err)
	epsilon, err := strconv.ParseInt(epsilonDigits, 2, 64)
	checkErr(err)

	fmt.Println("part 1: ", gamma*epsilon)
}

func readInput() []string {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
