package insertionSortAdvancedAnalysis

import "fmt"

func insertionSort(arr []int32) int32 {
	fmt.Println("before sort arr:", arr)

	sliceLength := int32(len(arr))
	var shifts int32
	var shifted bool

	for {
		shifted = false
		for i := int32(0); i < sliceLength-1; i++ {
			if arr[i] > arr[i+1] {
				shifted = true
				switchPositions(arr, i, i+1)
				shifts++
				fmt.Println("after switch arr:", arr)
			}
		}
		if !shifted {
			break
		}
	}

	fmt.Println("after sort arr:", arr)
	fmt.Println("shifts:", shifts)
	return shifts
}

func switchPositions(arr []int32, pos1 int32, pos2 int32) {
	var temp int32
	temp = arr[pos1]
	arr[pos1] = arr[pos2]
	arr[pos2] = temp
}
