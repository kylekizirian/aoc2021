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
	part2(template, rules)
}

func part1(template string, rules map[string]string) {
	current := template
	for i := 0; i < 10; i++ {
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

	min, max := minAndMax(ctr)
	fmt.Println("part1: ", max-min)
}

func part2(template string, rules map[string]string) {
	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		substr := string(template[i]) + string(template[i+1])
		pairs[substr]++
	}

	for i := 0; i < 40; i++ {
		nextPairs := make(map[string]int)
		for pair, count := range pairs {
			insertion := rules[pair]
			firstSubstr := string(pair[0]) + insertion
			secondSubstr := insertion + string(pair[1])
			nextPairs[firstSubstr] += count
			nextPairs[secondSubstr] += count
		}
		pairs = nextPairs
	}

	ctr := make(map[rune]int)
	for substr, count := range pairs {
		ctr[rune(substr[0])] += count
	}
	// account for last character in template string
	ctr[rune(template[len(template)-1])]++

	min, max := minAndMax(ctr)
	fmt.Println("part2: ", max-min)
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

// minAndMax returns the minimum and maximum values in ctr map
func minAndMax(ctr map[rune]int) (int, int) {
	min, max := math.MaxInt, math.MinInt
	for _, val := range ctr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
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
