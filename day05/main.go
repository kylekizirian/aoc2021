package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := readInput()
	part1(lines)
	part2(lines)
}

func part1(lines []Line) {
	// maps [x][y] to num points
	counter := make(map[int]map[int]int)
	for _, line := range lines {
		if !line.IsHorizontal() && !line.IsVertical() {
			continue
		}

		points := line.PointsOnLine()
		for _, point := range points {
			if counter[point[0]] == nil {
				counter[point[0]] = make(map[int]int)
			}
			counter[point[0]][point[1]]++
		}
	}

	var atLeast2Overlap int
	for _, xMap := range counter {
		for _, count := range xMap {
			if count >= 2 {
				atLeast2Overlap++
			}
		}
	}
	fmt.Println("part 1: ", atLeast2Overlap)
}

func part2(lines []Line) {
	// maps [x][y] to num points
	counter := make(map[int]map[int]int)
	for _, line := range lines {
		points := line.PointsOnLine()
		for _, point := range points {
			if counter[point[0]] == nil {
				counter[point[0]] = make(map[int]int)
			}
			counter[point[0]][point[1]]++
		}
	}

	var atLeast2Overlap int
	for _, xMap := range counter {
		for _, count := range xMap {
			if count >= 2 {
				atLeast2Overlap++
			}
		}
	}
	fmt.Println("part 2: ", atLeast2Overlap)
}

func readInput() []Line {
	var lines []Line
	numRe := regexp.MustCompile("\\d+")

	file, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		points := numRe.FindAllString(scanner.Text(), 4)
		line := Line{
			x1: strToInt(points[0]),
			y1: strToInt(points[1]),
			x2: strToInt(points[2]),
			y2: strToInt(points[3]),
		}
		lines = append(lines, line)
	}

	return lines
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (l Line) IsHorizontal() bool {
	return l.y1 == l.y2
}

func (l Line) IsVertical() bool {
	return l.x1 == l.x2
}

func (l Line) PointsOnLine() [][]int {

	if l.IsHorizontal() {
		var points [][]int
		start, end := l.x1, l.x2
		if l.x2 < start {
			start, end = l.x2, l.x1
		}
		for x := start; x <= end; x++ {
			points = append(points, []int{x, l.y1})
		}
		return points
	}

	if l.IsVertical() {
		var points [][]int
		start, end := l.y1, l.y2
		if l.y2 < start {
			start, end = l.y2, l.y1
		}
		for y := start; y <= end; y++ {
			points = append(points, []int{l.x1, y})
		}
		return points
	}

	var points [][]int
	startX, startY := l.x1, l.y1
	endX, endY := l.x2, l.y2
	if l.x2 < startX {
		startX, startY = l.x2, l.y2
		endX, endY = l.x1, l.y1
	}
	y := startY
	for x := startX; x <= endX; x++ {
		points = append(points, []int{x, y})
		if y < endY {
			y++
		} else {
			y--
		}
	}
	return points
}

func (l Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.x1, l.x2, l.y1, l.y2)
}

func strToInt(num string) int {
	val, err := strconv.Atoi(num)
	checkErr(err)
	return val
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
