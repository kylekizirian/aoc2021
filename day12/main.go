package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var startNode = Node{"start", false}

func main() {
	nodeMap := readInput()
	part1(nodeMap)
	part2(nodeMap)
}

func part1(nodeMap map[Node][]Node) {
	visitNext := func(next Node, visited []Node) bool {
		return next.big || !contains(visited, next)
	}
	isValid := func(visited []Node) bool { return true }
	numPaths := pathsToEnd(startNode, nil, nodeMap, visitNext, isValid)
	fmt.Println("part 1: ", numPaths)
}

func part2(nodeMap map[Node][]Node) {
	visitNext := func(next Node, visited []Node) bool {
		return next.big || !contains(visited, next) || !smallCaveVisitedTwice(visited)
	}
	numPaths := pathsToEnd(startNode, nil, nodeMap, visitNext, validPathPart2)
	fmt.Println("part 2: ", numPaths)
}

func pathsToEnd(curNode Node, visited []Node, nodeMap map[Node][]Node, visitNext func(Node, []Node) bool, isValidPath func([]Node) bool) int {
	if curNode.id == "end" {
		if isValidPath(visited) {
			return 1
		}
		return 0
	}

	paths := 0
	for _, next := range nodeMap[curNode] {
		if visitNext(next, visited) {
			visitedCopy := make([]Node, len(visited))
			copy(visitedCopy, visited)
			visitedCopy = append(visitedCopy, curNode)
			paths += pathsToEnd(next, visitedCopy, nodeMap, visitNext, isValidPath)
		}
	}
	return paths
}

type Node struct {
	id  string
	big bool
}

func contains(nodes []Node, node Node) bool {
	for _, n := range nodes {
		if n.id == node.id {
			return true
		}
	}
	return false
}

func smallCaveVisitedTwice(nodes []Node) bool {
	// check if we visit any small caves twice
	small := make(map[Node]bool)
	for _, node := range nodes {
		if !node.big {
			if small[node] {
				return true
			}
			small[node] = true
		}
	}
	return false
}

func validPathPart2(path []Node) bool {
	visitedTwice := false
	small := make(map[Node]bool)
	for _, node := range path {
		if !node.big {
			if small[node] {
				if visitedTwice {
					return false
				}
				visitedTwice = true
			}
			small[node] = true
		}
	}
	return true
}

func (n Node) String() string {
	return n.id
}

func readInput() map[Node][]Node {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	nodeMap := make(map[Node][]Node)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		node1 := Node{
			id:  parts[0],
			big: IsUpper(parts[0]),
		}
		node2 := Node{
			id:  parts[1],
			big: IsUpper(parts[1]),
		}
		if node1.id != "end" && node2.id != "start" {
			nodeMap[node1] = append(nodeMap[node1], node2)
		}
		if node2.id != "end" && node1.id != "start" {
			nodeMap[node2] = append(nodeMap[node2], node1)
		}
	}
	return nodeMap
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
