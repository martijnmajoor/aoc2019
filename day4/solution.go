package main

import (
	"fmt"
	"strconv"
)

func main() {
	min := 109165
	max := 576723

	fmt.Println("part one:", matches(min, max))
}
func matches(min int, max int) (matches int) {
	for i := min; i < max; i++ {
		pass := strconv.Itoa(i)
		if adjacent(pass) && !decreases(pass) {
			matches++
		}
	}
	return
}
func adjacent(pass string) bool {
	for idx := range pass {
		if idx > 0 && pass[idx] == pass[idx-1] {
			return true
		}
	}

	return false
}
func decreases(pass string) bool {
	for idx := range pass {
		if idx > 0 {
			current, _ := strconv.Atoi(string(pass[idx]))
			prev, _ := strconv.Atoi(string(pass[idx-1]))

			if current < prev {
				return true
			}
		}
	}

	return false
}
