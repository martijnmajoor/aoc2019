package main

import (
	"fmt"
	"strings"

	"github.com/martijnmajoor/aoc2018/utils/file"
)

func main() {
	orbits := file.ReadLines("orbits.txt")

	fmt.Println("part one:", count(orbits))
}
func count(orbits []string) (total int) {
	objects := direct(orbits)

	for _, other := range objects {
		total += follow(other, objects)
	}

	return
}
func direct(orbits []string) (objects map[string]string) {
	objects = make(map[string]string)

	for _, orbit := range orbits {
		direct := strings.Split(orbit, ")")
		objects[direct[1]] = direct[0]
	}

	return
}
func follow(target string, objects map[string]string) (steps int) {
	for target != "COM" {
		target = objects[target]
		steps++
	}

	steps++ // direct orbit to COM

	return
}
