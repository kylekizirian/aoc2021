package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var closeToOpen = map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}

func main() {
	lines := readInput()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	score := 0
	for _, line := range lines {
		score += corruptedScore(line)
	}
	fmt.Println("part 1: ", score)
}

func part2(lines []string) {
	var scores []int
	for _, line := range lines {
		score := completionScore(line)
		if score != 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Println("part 2: ", scores[len(scores)/2])
}

func corruptedScore(line string) int {
	charScore := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	var stack []string
	for _, char := range strings.Split(line, "") {
		switch char {
		case "(", "[", "{", "<":
			stack = append(stack, char)
		default:
			if stack[len(stack)-1] != closeToOpen[char] {
				return charScore[char]
			}
			stack = stack[:len(stack)-1]
		}
	}
	return 0
}

func completionScore(line string) int {
	charScore := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}
	var stack []string
	for _, char := range strings.Split(line, "") {
		switch char {
		case "(", "[", "{", "<":
			stack = append(stack, char)
		default:
			if stack[len(stack)-1] != closeToOpen[char] {
				return 0
			}
			stack = stack[:len(stack)-1]
		}
	}

	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		score += charScore[stack[i]]
	}

	return score
}

func readInput() []string {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

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
