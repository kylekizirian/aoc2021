package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	coords, folds := readInput()
	part1(coords, folds)
	part2(coords, folds)
}

func part1(coords CoordSet, folds []Fold) {
	firstFold := folds[0]
	foldedCoords := make(CoordSet)
	for coord := range coords {
		foldedCoords.Add(FoldCoord(coord, firstFold))
	}
	fmt.Println("part 1: ", len(foldedCoords))
}

func part2(coords CoordSet, folds []Fold) {
	for _, fold := range folds {
		foldedCoords := make(CoordSet)
		for coord := range coords {
			foldedCoords.Add(FoldCoord(coord, fold))
		}
		coords = foldedCoords
	}

	maxX, maxY := math.MinInt, math.MinInt
	for coord := range coords {
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}
	grid := make([][]string, maxY+1)
	for i := range grid {
		grid[i] = make([]string, maxX+1)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	for coord := range coords {
		grid[coord.Y][coord.X] = "*"
	}

	fmt.Println("part2:")
	for _, row := range grid {
		fmt.Println(row)
	}
}

type Fold struct {
	IsX  bool
	Line int
}

type Coord struct {
	X int
	Y int
}

func FoldCoord(coord Coord, fold Fold) Coord {
	if fold.IsX {
		if coord.X < fold.Line {
			return coord
		}
		dist := absInt(coord.X - fold.Line)
		return Coord{fold.Line - dist, coord.Y}
	}

	if coord.Y < fold.Line {
		return coord
	}
	dist := absInt(coord.Y - fold.Line)
	return Coord{coord.X, fold.Line - dist}
}

type CoordSet map[Coord]bool

func (c CoordSet) Add(coord Coord) {
	c[coord] = true
}

func (c CoordSet) Delete(coord Coord) {
	delete(c, coord)
}

func (c CoordSet) Has(coord Coord) bool {
	return c[coord]
}

func readInput() (CoordSet, []Fold) {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	coordSet := make(CoordSet)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ",")
		coordSet.Add(Coord{strToInt(parts[0]), strToInt(parts[1])})
	}

	var folds []Fold
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "fold") {
			continue
		}

		parts := strings.Fields(line)
		location := strings.Split(parts[2], "=")
		folds = append(folds, Fold{location[0] == "x", strToInt(location[1])})
	}

	return coordSet, folds
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
