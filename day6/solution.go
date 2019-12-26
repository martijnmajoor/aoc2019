package main

import (
	"fmt"
	"strings"

	"github.com/martijnmajoor/aoc2018/utils/file"
)

func main() {
	objects := direct(file.ReadLines("orbits.txt"))

	fmt.Println("part one:", sum(objects))
	fmt.Println("part two:", shortest(objects))
}
func direct(orbits []string) (objects map[string]string) {
	objects = make(map[string]string)

	for _, orbit := range orbits {
		direct := strings.Split(orbit, ")")
		objects[direct[1]] = direct[0]
	}

	return
}
func sum(objects map[string]string) (total int) {
	for object := range objects {
		total += len(transfers(object, "COM", objects))
	}

	return
}
func transfers(source string, target string, objects map[string]string) (transfers []string) {
	for source != target {
		transfers = append(transfers, objects[source])
		source = objects[source]
	}
	return
}
func shortest(objects map[string]string) int {
	you := transfers("YOU", "COM", objects)
	san := transfers("SAN", "COM", objects)
	intersect := intersect(you, san)

	return len(transfers("YOU", intersect, objects)) +
		len(transfers("SAN", intersect, objects)) - 2 // exclude steps to reach intersection
}
func intersect(you []string, san []string) string {
	for _, self := range you {
		for _, other := range san {
			if self == other {
				return self
			}
		}
	}
	return ""
}
