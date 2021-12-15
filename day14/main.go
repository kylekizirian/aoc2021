package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	template, rules := readInput()
	part1(template, rules)
}

func part1(template string, rules map[string]string) {
	current := template
	for i := 1; i < 11; i++ {
		next := ""
		for j := 0; j < len(current)-1; j++ {
			first := string(current[j])
			second := string(current[j+1])
			if insertion, ok := rules[first+second]; ok {
				next += first + insertion
			} else {
				next += first
			}
		}
		next += string(current[len(current)-1])
		current = next
	}

	ctr := make(map[rune]int)
	for _, r := range current {
		ctr[r]++
	}

	min, max := math.MaxInt, math.MinInt
	for _, val := range ctr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	fmt.Println("part1: ", max-min)
}

type Node struct {
	Char rune
	Next *Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	cur, str := n, ""
	for cur.Next != nil {
		str += string(n.Char) + "->"
		cur = n.Next
	}
	return str + string(n.Char)
}

func readInput() (string, map[string]string) {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	template := scanner.Text()

	// scan over the blank line between template and rules
	scanner.Scan()

	rules := make(map[string]string)
	for scanner.Scan() {
		rule := strings.Split(scanner.Text(), "->")
		rules[strings.TrimSpace(rule[0])] = strings.TrimSpace(rule[1])
	}

	return template, rules
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
