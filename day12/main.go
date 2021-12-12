package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	nodeMap := readInput()
	part1(nodeMap)
}

func part1(nodeMap map[Node][]Node) {
	fmt.Println("part 1: ", pathsToEnd(Node{"start", false}, nil, nodeMap))
}

func pathsToEnd(curNode Node, visited []Node, nodeMap map[Node][]Node) int {
	if curNode.id == "end" {
		return 1
	}

	paths := 0
	for _, next := range nodeMap[curNode] {
		// check if already visited
		if next.big || !contains(visited, next) {
			visitedCopy := make([]Node, len(visited))
			copy(visitedCopy, visited)
			visitedCopy = append(visitedCopy, curNode)
			paths += pathsToEnd(next, visitedCopy, nodeMap)
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
