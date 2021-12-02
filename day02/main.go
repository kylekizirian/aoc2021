package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	movements := readInput()
	partA(movements)
	partB(movements)
}

func partA(movements []string) {
	var horiz, depth int

	for _, movement := range movements {
		tokens := strings.Fields(movement)
		direction := tokens[0]
		amount, err := strconv.Atoi(tokens[1])
		checkErr(err)

		switch direction {
		case "forward":
			horiz += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}

	fmt.Println("part A: ", horiz*depth)
}

func partB(movements []string) {
	var aim, horiz, depth int

	for _, movement := range movements {
		tokens := strings.Fields(movement)
		direction := tokens[0]
		amount, err := strconv.Atoi(tokens[1])
		checkErr(err)

		switch direction {
		case "forward":
			horiz += amount
			depth += amount * aim
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}
	fmt.Println("part B: ", horiz*depth)
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
