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
	timers := readInput()
	part1(timers)
}

func part1(timers []int) {
	timerHeap := make(IntHeap, len(timers))
	for i, timer := range timers {
		timerHeap[i] = timer
	}
	heap.Init(&timerHeap)

	for timerHeap[0] < 80 {
		next := heap.Pop(&timerHeap).(int)
		heap.Push(&timerHeap, next+7)
		heap.Push(&timerHeap, next+9)
	}
	fmt.Println("part 1: ", len(timerHeap))
}

func readInput() []int {

	file, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fields := strings.Split(scanner.Text(), ",")
	timers := make([]int, len(fields))

	for i, field := range fields {
		timers[i] = strToInt(field)
	}

	return timers
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
