package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	program := read("program.txt")
	noun, verb := find(19690720, program)

	fmt.Println("part one:", run(prep(program, 12, 2)))
	fmt.Println("part two:", 100*noun+verb)
}
func read(file string) (program []int) {
	if b, err := ioutil.ReadFile(file); err == nil {
		for _, s := range strings.Split(string(b), ",") {
			code, _ := strconv.Atoi(s)
			program = append(program, code)
		}
	}
	return
}
func run(p []int) int {
	step := 0
	halt := 99
	quit := false

	for !quit {
		switch p[step] {
		case 1:
			p[p[step+3]] = p[p[step+1]] + p[p[step+2]]
		case 2:
			p[p[step+3]] = p[p[step+1]] * p[p[step+2]]
		}
		step += 4
		quit = p[step] == halt
	}

	return p[0]
}
func prep(src []int, noun int, verb int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	dest[1] = noun
	dest[2] = verb
	return dest
}
func find(output int, program []int) (noun int, verb int) {
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			if run(prep(program, noun, verb)) == output {
				return
			}
		}
	}
	return -1, -1
}
