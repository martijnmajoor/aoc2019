package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	min := 109165
	max := 576723

	fmt.Println("part one:", matches(min, max, false))
	fmt.Println("part two:", matches(min, max, true))
}
func matches(min int, max int, exact bool) (matches int) {
	for i := min; i <= max; i++ {
		pass := strconv.Itoa(i)
		if adjacent(pass, exact) && !decreases(pass) {
			matches++
		}
	}
	return
}
func adjacent(pass string, exact bool) bool {
	for idx := range pass {
		if idx < len(pass)-1 && pass[idx] == pass[idx+1] {
			if !exact || !larger(pass, pass[idx]) {
				return true
			}
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
func larger(pass string, digit byte) bool {
	pattern := fmt.Sprintf("%v{3,}", string(digit))
	match, _ := regexp.MatchString(pattern, pass)

	return match
}
