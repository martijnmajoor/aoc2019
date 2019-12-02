package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/martijnmajoor/aoc2018/utils/file"
)

func main() {
	masses := file.ReadLines("module_masses.txt")

	fmt.Println("part one:", int(fuelTotal(masses)))
	fmt.Println("part two:", int(fuelInclusiveTotal(masses)))
}
func fuelTotal(masses []string) (total float64) {
	for _, mass := range masses {
		f, _ := strconv.ParseFloat(mass, 64)
		total += calcFuel(f)
	}
	return
}
func calcFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2
}
func fuelInclusiveTotal(masses []string) (total float64) {
	for _, mass := range masses {
		f, _ := strconv.ParseFloat(mass, 64)
		total += calcFuelInclusive(f)
	}
	return
}
func calcFuelInclusive(fuel float64) (total float64) {
	sub := calcFuel(fuel)
	for sub > 0 {
		total += sub
		sub = calcFuel(sub)
	}
	return
}
