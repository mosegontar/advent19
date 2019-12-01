package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// CalculateFuel returns fuel required
// to launch a given module based on the
// the module's mass.
func CalculateFuel(mass int) int {
	return (mass / 3) - 2
}

func CalculateFuelRecurse(mass int) int {
	fuel := CalculateFuel(mass)
	if fuel <= 0 {
		return 0
	} else {
		return fuel + CalculateFuelRecurse(fuel)
	}
}

func ReadMasses(reader io.Reader) []int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	masses := []int{}
	for scanner.Scan() {
		masses = append(masses, convertToInt(scanner.Text()))
	}

	return masses
}

func SumInts(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func convertToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic("Couldn't convert")
	}
	return val
}

func SumFuelReqs(masses []int, fuelCalc func(mass int) int) int {
	fuelReqs := []int{}
	for i := 0; i < len(masses); i++ {
		fuelReqs = append(fuelReqs, fuelCalc(masses[i]))
	}
	return SumInts(fuelReqs)
}

func day1part1(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	masses := ReadMasses(f)
	fmt.Printf("Day 1 Part 1: %d\n", SumFuelReqs(masses, CalculateFuel))
}

func day1part2(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	masses := ReadMasses(f)
	fmt.Printf("Day 1 Part 1: %d\n", SumFuelReqs(masses, CalculateFuelRecurse))
}

func main() {
	day1part1("day1/input1.txt")
	day1part2("day1/input1.txt")
}
