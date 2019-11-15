package insertionSortAdvancedAnalysis

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {

	slice := []int32{1, 1, 1, 2, 2}

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
