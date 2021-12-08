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
	// split on | delimeter
	// count num words w/ len 1, 3, 4 and 7

	var outputs []string
	for _, line := range lines {
		fields := strings.Split(line, "|")
		outs := strings.Fields(fields[1])
		outputs = append(outputs, outs...)
	}

	count := 0
	for _, output := range outputs {
		switch len(output) {
		case 1, 2, 3, 4, 7:
			count++
		}
	}

	fmt.Println(count)
}

func readInput() []string {
	file, err := os.Open("./input.txt")
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
