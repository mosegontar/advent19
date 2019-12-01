package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	got := CalculateFuel(100756)
	want := 33583

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCalculateFuelRecurse(t *testing.T) {
	got := CalculateFuelRecurse(100756)
	want := 50346

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadMasses(t *testing.T) {
	massData := []string{"1", "2", "34", "456", "7890"}
	massReader := strings.NewReader(strings.Join(massData, "\n"))

	got := ReadMasses(massReader)
	want := []int{1, 2, 34, 456, 7890}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumInts(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	got := SumInts(ints)
	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumFuelReqs(t *testing.T) {
	masses := []int{12, 14, 1969, 100756}
	got := SumFuelReqs(masses, CalculateFuel)
	want := 34241

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumFuelReqsRecursive(t *testing.T) {
	masses := []int{14, 1969, 100756}
	got := SumFuelReqs(masses, CalculateFuelRecurse)
	want := 51314

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
