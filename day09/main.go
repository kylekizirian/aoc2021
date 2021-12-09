package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	matrix := readInput()
	part1(matrix)
	part2(matrix)
}

func part1(matrix Grid) {
	totalRisk := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix.IsLowPoint(Location{row, col}) {
				totalRisk += matrix[row][col] + 1
			}
		}
	}
	fmt.Println("part 1: ", totalRisk)
}

func part2(grid Grid) {
	// 9s cannot be part of basins
	// we could do one pass and mark all low points as basins with unique IDs,
	// then BFS from low point, marking all points with the basin ID until we
	// 9s or are surrounded by other "marked" locations

	basinIdCtr := 1
	basinIds := make([][]int, len(grid))
	for i := range grid {
		basinIds[i] = make([]int, len(grid[i]))
	}

	var stack []Location
	basinIdSize := make(map[int]int)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid.IsLowPoint(Location{row, col}) {
				basinIds[row][col] = basinIdCtr
				basinIdSize[basinIdCtr] = 1
				basinIdCtr++
				stack = append(stack, Location{row, col})
			}
		}
	}

	for len(stack) > 0 {
		loc := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		basinId := basinIds[loc.Row][loc.Col]
		if basinId == 0 {
			panic("basin id of 0")
		}

		// look at all surrounding
		for _, adj := range grid.Adjacent(loc) {
			if grid[adj.Row][adj.Col] != 9 && basinIds[adj.Row][adj.Col] == 0 {
				basinIds[adj.Row][adj.Col] = basinId
				basinIdSize[basinId]++
				stack = append(stack, adj)
			}
		}
	}

	var sizes []int
	for _, size := range basinIdSize {
		sizes = append(sizes, size)
	}

	sort.Ints(sizes)
	l := len(sizes)

	fmt.Println("part 2: ", sizes[l-3]*sizes[l-2]*sizes[l-1])
}

func readInput() Grid {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var grid Grid
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "")
		nums := make([]int, len(vals))
		for i, val := range vals {
			nums[i] = strToInt(val)
		}
		grid = append(grid, nums)
	}

	return grid
}

type Location struct {
	Row int
	Col int
}

type Grid [][]int

func (g Grid) IsLowPoint(loc Location) bool {
	for _, adj := range g.Adjacent(loc) {
		if g[adj.Row][adj.Col] <= g[loc.Row][loc.Col] {
			return false
		}
	}
	return true
}

func (g Grid) Adjacent(loc Location) []Location {
	var adjacent []Location

	if loc.Col > 0 {
		adjacent = append(adjacent, Location{loc.Row, loc.Col - 1})
	}
	if loc.Row > 0 {
		adjacent = append(adjacent, Location{loc.Row - 1, loc.Col})
	}
	if loc.Col < len(g[loc.Row])-1 {
		adjacent = append(adjacent, Location{loc.Row, loc.Col + 1})
	}
	if loc.Row < len(g)-1 {
		adjacent = append(adjacent, Location{loc.Row + 1, loc.Col})
	}

	return adjacent
}

func strToInt(val string) int {
	num, err := strconv.Atoi(val)
	checkErr(err)
	return num
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
