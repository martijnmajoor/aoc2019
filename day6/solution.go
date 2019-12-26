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

	for object := range objects {
		total += len(transfers(object, "COM", objects))
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
func transfers(source string, target string, objects map[string]string) (transfers []string) {
	for source != target {
		transfers = append(transfers, objects[source])
		source = objects[source]
	}
	return
}
