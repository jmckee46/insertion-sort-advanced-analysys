package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func insertionSort(arr []int32) int64 {

	a := int64(0)
	var shifts = &a
	_ = mergeSort(arr, shifts)

	return *shifts
}

func mergeSort(arr []int32, shifts *int64) []int32 {
	arrLength := int32(len(arr))

	if arrLength == 1 {
		return arr
	}

	arrMiddle := arrLength / 2
	var left = make([]int32, arrMiddle)
	var right = make([]int32, arrLength-arrMiddle)

	for i := int32(0); i < arrLength; i++ {
		if i < arrMiddle {
			left[i] = arr[i]
		} else {
			right[i-arrMiddle] = arr[i]
		}
	}

	return merge(mergeSort(left, shifts), mergeSort(right, shifts), shifts)
}

func merge(left, right []int32, shifts *int64) []int32 {
	result := make([]int32, len(left)+len(right))

	i := int32(0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
			*shifts += int64(len(left))
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

// this version needs to change shifts to int64
// func insertionSort(arr []int32) int32 {

// 	sliceLength := int32(len(arr))
// 	var shifts int32

// 	for i := int32(0); i < sliceLength; i++ {
// 		for j := i + 1; j < sliceLength; j++ {
// 			if arr[i] > arr[j] {
// 				shifts++
// 			}
// 		}
// 	}

// 	return shifts
// }

// this version needs to change shifts to int64
// func insertionSort(arr []int32) int32 {

// 	sliceLength := int32(len(arr))
// 	var shifts int32
// 	var temp int32

// 	for i := int32(0); i < sliceLength-1; i++ {
// 		if arr[i] > arr[i+1] {
// 			temp = arr[i]
// 			arr[i] = arr[i+1]
// 			arr[i+1] = temp
// 			shifts++
// 			for j := i; j > 0; j-- {
// 				if arr[j-1] > arr[j] {
// 					temp = arr[j-1]
// 					arr[j-1] = arr[j]
// 					arr[j] = temp
// 					shifts++
// 				}
// 			}
// 		}
// 	}

// 	return shifts
// }

func main() {
	file, err := os.Open("test-case-7")
	checkError(err)

	reader := bufio.NewReaderSize(file, 1024*1024)

	stdout, err := os.Create("OUTPUT-test-case-7")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := insertionSort(arr)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
