package main

import (
	"fmt"

	"github.com/martijnmajoor/aoc2019/program"
)

func main() {
	opcodes := program.Read("program.txt")

	fmt.Println("part one:", program.Diagnose(opcodes, 1))
}
