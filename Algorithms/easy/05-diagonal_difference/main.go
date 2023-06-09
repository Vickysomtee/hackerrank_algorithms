package main

import (
	"fmt"
)

var arr = [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}

func diagonalDifference(arr [][]int) int {
	diagonalA := 0
	diagonalB := 0

	for i := 0; i < len(arr); i++ {
		diagonalA += arr[i][i]
		diagonalB += arr[i][len(arr)-i-1]
	}

	value := diagonalA - diagonalB

	// Added this if statement to get the absolute value
	// since Golang's inbuilt Math.Abs works only float number types
	if value < 0 {
		value = -value
	}

	return value
}

func main() {
	result := diagonalDifference(arr)

	fmt.Println(result)
}
