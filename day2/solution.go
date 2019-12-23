package main

import (
	"fmt"

	"github.com/martijnmajoor/aoc2019/program"
)

func main() {
	opcodes := program.Read("program.txt")
	noun, verb := find(19690720, opcodes)

	fmt.Println("part one:", program.Run(prep(opcodes, 12, 2)))
	fmt.Println("part two:", 100*noun+verb)
}
func prep(src []int, noun int, verb int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	dest[1] = noun
	dest[2] = verb
	return dest
}
func find(output int, codes []int) (noun int, verb int) {
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			if program.Run(prep(codes, noun, verb)) == output {
				return
			}
		}
	}
	return -1, -1
}
