package main

import (
	"fmt"

	"github.com/martijnmajoor/aoc2019/program"
)

func main() {
	fmt.Println("part one:", program.Diagnose(program.Read("program.txt"), 1))
	fmt.Println("part two:", program.Diagnose(program.Read("program.txt"), 5))
}
