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
	part2(lines)
}

func part1(binNums []string) {
	var gammaDigits, epsilonDigits string
	for pos := range binNums[0] {
		switch mostCommonBit(binNums, pos) {
		case 0:
			gammaDigits += "0"
			epsilonDigits += "1"
		case 1:
			gammaDigits += "1"
			epsilonDigits += "0"
		}
	}

	gamma, err := strconv.ParseInt(gammaDigits, 2, 64)
	checkErr(err)
	epsilon, err := strconv.ParseInt(epsilonDigits, 2, 64)
	checkErr(err)

	fmt.Println("part 1: ", gamma*epsilon)
}

func part2(binNums []string) {
	oxygenNums := binNums
	co2Nums := binNums

	pos := 0
	for len(oxygenNums) > 1 {
		switch mostCommonBit(oxygenNums, pos) {
		case 0:
			oxygenNums = filter(oxygenNums, "0", pos)
		case -1, 1:
			oxygenNums = filter(oxygenNums, "1", pos)
		}
		pos++
	}

	pos = 0
	for len(co2Nums) > 1 {
		switch mostCommonBit(co2Nums, pos) {
		case 0:
			co2Nums = filter(co2Nums, "1", pos)
		case -1, 1:
			co2Nums = filter(co2Nums, "0", pos)
		}
		pos++
	}

	oxygen, err := strconv.ParseInt(oxygenNums[0], 2, 64)
	checkErr(err)
	co2, err := strconv.ParseInt(co2Nums[0], 2, 64)
	checkErr(err)

	fmt.Println("part 2: ", oxygen*co2)
}

// mostCommonBit loops over all bitStrings and returns 0 if it's the most common bit,
// or 1 if it's the most common bit. returns -1 if they are equally common
func mostCommonBit(bitStrings []string, pos int) int {
	// increment when we see a 1, decrement on 0
	counter := 0
	for _, binNum := range bitStrings {
		switch string(binNum[pos]) {
		case "0":
			counter -= 1
		case "1":
			counter += 1
		default:
			panic("unexpected digit in binary string")
		}
	}

	if counter == 0 {
		return -1
	} else if counter < 0 {
		return 0
	} else {
		return 1
	}
}

// filter returns all strings with given bit in pos
func filter(bitStrings []string, bit string, pos int) []string {
	var res []string
	for _, bitString := range bitStrings {
		if string(bitString[pos]) == bit {
			res = append(res, bitString)
		}
	}
	return res
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
