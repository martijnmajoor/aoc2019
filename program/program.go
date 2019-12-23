package program

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Read(file string) (program []int) {
	if b, err := ioutil.ReadFile(file); err == nil {
		for _, s := range strings.Split(string(b), ",") {
			code, _ := strconv.Atoi(s)
			program = append(program, code)
		}
	}
	return
}

func Run(codes []int) int {
	code, _ := process(codes, 0)

	return code
}
func Diagnose(codes []int, input int) int {
	_, output := process(codes, input)

	return output
}
func process(codes []int, input int) (int, int) {
	step := 0
	halt := 99
	quit := false
	outputs := []int{-1}

	for !quit {
		code := parse(codes, step)

		switch code[0] {
		case 1:
			codes[code[3]] = code[1] + code[2]
			step += 4
		case 2:
			codes[code[3]] = code[1] * code[2]
			step += 4
		case 3:
			codes[code[1]] = input
			step += 2
		case 4:
			outputs = append(outputs, codes[code[1]])
			step += 2
		case 5:
			step += 3
			if code[1] != 0 {
				step = code[2]
			}
		case 6:
			step += 3
			if code[1] == 0 {
				step = code[2]
			}
		case 7:
			codes[code[3]] = cond(code[1] < code[2])
			step += 4
		case 8:
			codes[code[3]] = cond(code[1] == code[2])
			step += 4
		}

		quit = codes[step] == halt
	}

	return codes[0], outputs[len(outputs)-1]
}
func parse(codes []int, idx int) []int {
	// inspired by https://github.com/Ullaakut/aoc19
	op := codes[idx] % 100
	left := get(codes, idx+1, immediate(op, (codes[idx]%1000)/100))
	right := get(codes, idx+2, immediate(op, (codes[idx]%10000)/1000))

	return []int{op, left, right, codes[idx+3]}
}
func get(codes []int, idx int, immediate bool) int {
	if immediate {
		return codes[idx]
	}

	return codes[codes[idx]]
}
func immediate(op int, code int) bool {
	return op == 3 || op == 4 || code == 1
}
func cond(validates bool) int {
	if validates {
		return 1
	}

	return 0
}
