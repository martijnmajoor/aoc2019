package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnmajoor/aoc2018/utils/file"
)

func main() {
	wires := file.ReadLines("wires.txt")
	nearest, shortest := follow(wires)

	fmt.Println("part one:", nearest)
	fmt.Println("part two:", shortest)
}
func follow(wires []string) (nearest int, shortest int) {
	grid := make(map[int]map[int]int)
	nearest = 0
	shortest = 0

	for idx, wire := range wires {
		moves := strings.Split(wire, ",")
		x := 0
		y := 0
		steps := 0

		for _, move := range moves {
			amount, _ := strconv.Atoi(move[1:])

			for i := 0; i < amount; i++ {
				steps++

				switch move[0] {
				case 'U':
					y++
				case 'D':
					y--
				case 'R':
					x++
				case 'L':
					x--
				}

				if idx == 0 {
					place(grid, x, y, idx, steps)
				} else if _, ok := grid[x][y]; ok {
					nearest = min(dist(x, y), nearest)
					shortest = min(grid[x][y]+steps, shortest)
				}
			}
		}
	}

	return
}
func min(self int, other int) int {
	if other == 0 || self < other {
		return self
	}

	return other
}
func dist(x int, y int) int {
	return abs(x) + abs(y)
}
func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
func place(grid map[int]map[int]int, x int, y int, wire int, steps int) {
	if _, ok := grid[x]; !ok {
		grid[x] = make(map[int]int)
	}

	grid[x][y] = min(steps, grid[x][y])
}
