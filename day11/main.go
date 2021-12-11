package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := readInput()
	part1(grid)
}

func part1(grid Grid) {
	numFlashed := 0
	for i := 0; i < 100; i++ {
		numFlashed += round(grid)
	}
	fmt.Println("part 1: ", numFlashed)
}

func round(grid Grid) int {
	// increase all by 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j]++
		}
	}

	hasFlashed := make([][]bool, len(grid))
	for i := range hasFlashed {
		hasFlashed[i] = make([]bool, len(grid[i]))
	}

	numFlashed := 0
	sigue := true
	for sigue {
		sigue = false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] > 9 && !hasFlashed[i][j] {
					sigue = true
					hasFlashed[i][j] = true
					numFlashed++
					for _, adj := range grid.Adjacent(Location{i, j}) {
						grid[adj.Row][adj.Col]++
					}
				}
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 9 {
				grid[i][j] = 0
			}
		}
	}

	return numFlashed
}

type Location struct {
	Row int
	Col int
}

type Grid [][]int

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
	if loc.Row > 0 && loc.Col > 0 {
		adjacent = append(adjacent, Location{loc.Row - 1, loc.Col - 1})
	}
	if loc.Row > 0 && loc.Col < len(g[loc.Row])-1 {
		adjacent = append(adjacent, Location{loc.Row - 1, loc.Col + 1})
	}
	if loc.Row < len(g)-1 && loc.Col > 0 {
		adjacent = append(adjacent, Location{loc.Row + 1, loc.Col - 1})
	}
	if loc.Row < len(g)-1 && loc.Col < len(g[loc.Row])-1 {
		adjacent = append(adjacent, Location{loc.Row + 1, loc.Col + 1})
	}

	return adjacent
}

func readInput() Grid {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var grid Grid
	for scanner.Scan() {
		var line []int
		for _, num := range strings.Split(scanner.Text(), "") {
			line = append(line, strToInt(num))
		}
		grid = append(grid, line)
	}
	return grid
}

func strToInt(s string) int {
	n, err := strconv.Atoi(s)
	checkErr(err)
	return n
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
