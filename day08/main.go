package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	TOP = iota
	TOP_LEFT
	TOP_RIGHT
	MIDDLE
	BOTTOM_RIGHT
	BOTTOM_LEFT
	BOTTOM
)

func main() {
	entries := readInput()
	part1(entries)
	part2(entries)
}

func part1(entries []Entry) {

	var outputs []string
	for _, entry := range entries {
		outputs = append(outputs, entry.output...)
	}

	count := 0
	for _, output := range outputs {
		switch len(output) {
		case 2, 3, 4, 7:
			count++
		}
	}

	fmt.Println("part 1: ", count)
}

func part2(entries []Entry) {
	sum := 0
	for _, entry := range entries {
		sum += Decode(entry)
	}
	fmt.Println("part 2: ", sum)
}

// Given an entry, return the number that its output represents
func Decode(entry Entry) int {
	lenToPatterns := make(map[int][]string)
	posToChar := make(map[int]string)

	for _, pattern := range entry.patterns {
		lenToPatterns[len(pattern)] = append(lenToPatterns[len(pattern)], pattern)
	}

	// take char in length 3 but not in length 2 to be top
	for _, char := range strings.Split(lenToPatterns[3][0], "") {
		if !strings.Contains(lenToPatterns[2][0], char) {
			posToChar[TOP] = char
		}
	}

	// get chars in len 4 but not in len 2, these are candidates for top left and middle
	var topLeftOrMiddle []string
	for _, char := range strings.Split(lenToPatterns[4][0], "") {
		if !strings.Contains(lenToPatterns[2][0], char) {
			topLeftOrMiddle = append(topLeftOrMiddle, char)
		}
	}

	// figure out top right, bottom right, middle, and top left
	for _, lenSixPattern := range lenToPatterns[6] {

		twoChars := strings.Split(lenToPatterns[2][0], "")
		if strings.Contains(lenSixPattern, twoChars[0]) && !strings.Contains(lenSixPattern, twoChars[1]) {
			posToChar[BOTTOM_RIGHT] = twoChars[0]
			posToChar[TOP_RIGHT] = twoChars[1]
		}
		if !strings.Contains(lenSixPattern, twoChars[0]) && strings.Contains(lenSixPattern, twoChars[1]) {
			posToChar[BOTTOM_RIGHT] = twoChars[1]
			posToChar[TOP_RIGHT] = twoChars[0]
		}

		if strings.Contains(lenSixPattern, topLeftOrMiddle[0]) && !strings.Contains(lenSixPattern, topLeftOrMiddle[1]) {
			posToChar[TOP_LEFT] = topLeftOrMiddle[0]
			posToChar[MIDDLE] = topLeftOrMiddle[1]
		}
		if !strings.Contains(lenSixPattern, topLeftOrMiddle[0]) && strings.Contains(lenSixPattern, topLeftOrMiddle[1]) {
			posToChar[TOP_LEFT] = topLeftOrMiddle[1]
			posToChar[MIDDLE] = topLeftOrMiddle[0]
		}
	}

	// find 2 remaining digits
	var remaining []string
	for _, char := range strings.Split("abcdefg", "") {
		switch char {
		case posToChar[TOP], posToChar[TOP_LEFT], posToChar[TOP_RIGHT], posToChar[MIDDLE], posToChar[BOTTOM_RIGHT]:
		default:
			remaining = append(remaining, char)
		}
	}

	// if we find a length 5 or length 6 pattern with 1 remaining char but not the other, that must be bottom
	fiveOrSixLens := append(lenToPatterns[5], lenToPatterns[6]...)
	for _, pattern := range fiveOrSixLens {
		if strings.Contains(pattern, remaining[0]) && !strings.Contains(pattern, remaining[1]) {
			posToChar[BOTTOM] = remaining[0]
			posToChar[BOTTOM_LEFT] = remaining[1]
			break
		}
		if !strings.Contains(pattern, remaining[0]) && strings.Contains(pattern, remaining[1]) {
			posToChar[BOTTOM] = remaining[1]
			posToChar[BOTTOM_LEFT] = remaining[0]
			break
		}
	}

	decoded := 0
	for i, output := range entry.output {
		mult := 1
		switch i {
		case 0:
			mult = 1000
		case 1:
			mult = 100
		case 2:
			mult = 10
		}

		switch len(output) {
		case 2: // must be 1
			decoded += 1 * mult
		case 3: // must be 7
			decoded += 7 * mult
		case 4: // must be 4
			decoded += 4 * mult
		case 5: // could be 2, 3, 5
			if !strings.Contains(output, posToChar[TOP_LEFT]) && !strings.Contains(output, posToChar[BOTTOM_RIGHT]) {
				decoded += 2 * mult
			} else if !strings.Contains(output, posToChar[TOP_LEFT]) && !strings.Contains(output, posToChar[BOTTOM_LEFT]) {
				decoded += 3 * mult
			} else {
				decoded += 5 * mult
			}
		case 6: // could be 0, 6, 9
			if !strings.Contains(output, posToChar[TOP_RIGHT]) {
				decoded += 6 * mult
			} else if !strings.Contains(output, posToChar[BOTTOM_LEFT]) {
				decoded += 9 * mult
			}
		case 7: // must be 8
			decoded += 8 * mult
		}
	}

	return decoded
}

type Entry struct {
	patterns []string
	output   []string
}

func readInput() []Entry {
	file, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var entries []Entry
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		entry := Entry{
			strings.Fields(line[0]),
			strings.Fields(line[1]),
		}
		entries = append(entries, entry)
	}

	return entries
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
