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

func SumFuelReqs(masses []int) int {
	fuelReqs := []int{}
	for i := 0; i < len(masses); i++ {
		fuelReqs = append(fuelReqs, CalculateFuel(masses[i]))
	}
	return SumInts(fuelReqs)
}

func main() {
	f, err := os.Open("day1/input1.txt")
	if err != nil {
		panic(err)
	}
	masses := ReadMasses(f)
	fmt.Println(SumFuelReqs(masses))
}
