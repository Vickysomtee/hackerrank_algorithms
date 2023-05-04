package main

import (
	"fmt"
	"sort"
)

// Test case 1
var numbers = []int{7, 69, 2, 221, 8974}

// Test case 2
// var numbers = []int{1, 2, 3, 4, 5}

func minMaxSum(arr []int) {

	min := 0
	max := 0

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		if i < len(arr)-1 {
			min += arr[i]
		}

		if i > 0 {
			max += arr[i]
		}
	}

	fmt.Println(min, max)
}

func main() {
	minMaxSum(numbers)
}
