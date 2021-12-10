package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readInput()
	part1(lines)
}

func part1(lines []string) {
	score := 0
	for _, line := range lines {
		score += corruptedScore(line)
	}
	fmt.Println("part 1: ", score)
}

func corruptedScore(line string) int {
	var stack []string
	for _, char := range strings.Split(line, "") {
		switch char {
		case "(", "[", "{", "<":
			stack = append(stack, char)
		case ")":
			if stack[len(stack)-1] != "(" {
				return 3
			}
			stack = stack[:len(stack)-1]
		case "]":
			if stack[len(stack)-1] != "[" {
				return 57
			}
			stack = stack[:len(stack)-1]
		case "}":
			if stack[len(stack)-1] != "{" {
				return 1197
			}
			stack = stack[:len(stack)-1]
		case ">":
			if stack[len(stack)-1] != "<" {
				return 25137
			}
			stack = stack[:len(stack)-1]
		}
	}
	return 0
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
