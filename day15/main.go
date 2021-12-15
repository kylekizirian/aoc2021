package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := readInput()
	part1(grid)
	part2(grid)
}

func part1(grid [][]int) {
	fmt.Println("part 1: ", minRisk(grid))
}

func part2(grid [][]int) {
	// we'll need to repeat the grid 8 times and append 4 times to the right,
	// then append 4 times below
	tiles := make([][][]int, 9)
	tiles[0] = grid
	for i := 1; i < 9; i++ {
		if i == 0 {
			tiles[i] = repeatGrid(grid)
		} else {
			tiles[i] = repeatGrid(tiles[i-1])
		}
	}

	rows := make([][][]int, 5)
	for i := range rows {
		row := tiles[i]
		for j := i + 1; j < i+5; j++ {
			for r := 0; r < len(row); r++ {
				row[r] = append(row[r], tiles[j][r]...)
			}
		}
		rows[i] = row
	}

	var fullGrid [][]int
	for _, row := range rows {
		fullGrid = append(fullGrid, row...)
	}

	// fmt.Println("len fullgrid: ", len(fullGrid))
	// fmt.Println("len fullgrid[0]: ", len(fullGrid[0]))
	// for _, r := range fullGrid {
	// 	fmt.Println(r)
	// }

	fmt.Println("part 2: ", minRisk(fullGrid))
}

// repeatGrid creates a copy of grid with each value incremented by 1,
// wrapping values of 9 back around to 1
func repeatGrid(grid [][]int) [][]int {
	repeated := make([][]int, len(grid))
	for i := range repeated {
		repeated[i] = make([]int, len(grid[i]))
		for j := 0; j < len(repeated[i]); j++ {
			repeated[i][j] = grid[i][j] + 1
			if repeated[i][j] > 9 {
				repeated[i][j] = 1
			}
		}
	}

	return repeated
}

type Position struct {
	Row, Col int
}

func (p Position) DecRow() Position {
	return Position{p.Row - 1, p.Col}
}

func (p Position) DecCol() Position {
	return Position{p.Row, p.Col - 1}
}

func (p Position) IncRow() Position {
	return Position{p.Row + 1, p.Col}
}

func (p Position) IncCol() Position {
	return Position{p.Row, p.Col + 1}
}

type Grid [][]int

func (g Grid) Adjacent(pos Position) []Position {
	var adjacent []Position
	if pos.Row > 0 {
		adjacent = append(adjacent, Position{pos.Row - 1, pos.Col})
	}
	if pos.Col > 0 {
		adjacent = append(adjacent, Position{pos.Row, pos.Col - 1})
	}
	if pos.Row < len(g)-1 {
		adjacent = append(adjacent, Position{pos.Row + 1, pos.Col})
	}
	if pos.Col < len(g[pos.Row])-1 {
		adjacent = append(adjacent, Position{pos.Row, pos.Col + 1})
	}
	return adjacent
}

type Risk struct {
	Pos     Position
	MinRisk int
}

type RiskHeap []Risk

func (h RiskHeap) Len() int           { return len(h) }
func (h RiskHeap) Less(i, j int) bool { return h[i].MinRisk < h[j].MinRisk }
func (h RiskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *RiskHeap) Push(x interface{}) {
	*h = append(*h, x.(Risk))
}

func (h *RiskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minRisk(grid Grid) int {
	h := &RiskHeap{Risk{Position{0, 0}, 0}}
	heap.Init(h)
	visited := make(map[Position]bool)
	dest := Position{len(grid) - 1, len(grid[0]) - 1}

	for {
		riskPos := heap.Pop(h).(Risk)
		if visited[riskPos.Pos] {
			continue
		}
		visited[riskPos.Pos] = true

		if riskPos.Pos == dest {
			return riskPos.MinRisk
		}

		for _, adj := range grid.Adjacent(riskPos.Pos) {
			if visited[adj] {
				continue
			}

			heap.Push(h, Risk{adj, riskPos.MinRisk + grid[adj.Row][adj.Col]})
		}
	}
}

func readInput() [][]int {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var output [][]int

	for scanner.Scan() {
		var line []int
		for _, val := range strings.Split(scanner.Text(), "") {
			line = append(line, strToInt(val))
		}
		output = append(output, line)
	}

	return output
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
