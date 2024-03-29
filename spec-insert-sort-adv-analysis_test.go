package main

// package insertionSortAdvancedAnalysis

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	// slice := []int32{1, 1, 1, 2, 2}
	slice := []int32{1, 1, 1, 2, 2, 3, 4, 5, 6}

	shifts := insertionSort(slice)
	if shifts != 0 {
		t.Errorf("got %d shifts instead of 0", shifts)
	}

	slice = []int32{2, 1, 3, 1, 2}

	shifts = insertionSort(slice)
	if shifts != 4 {
		t.Errorf("got %d shifts instead of 4", shifts)
	}

	slice = []int32{4, 3, 2, 1}

	shifts = insertionSort(slice)
	if shifts != 6 {
		t.Errorf("got %d shifts instead of 6", shifts)
	}
}

func TestInsertionSortLong(t *testing.T) {

	slice := make([]int32, 100000)

	for i := int32(0); i < 100000; i++ {
		slice[i] = 100000 - i
	}

	shifts := insertionSort(slice)
	if shifts != 704982704 {
		t.Errorf("got %d shifts instead of 704982704", shifts)
	}
	fmt.Println("1st run was 9.011s")
	fmt.Println("counting inversions was 4.213s")
}
