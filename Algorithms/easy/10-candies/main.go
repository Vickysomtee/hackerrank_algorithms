package main

import (
	"fmt"
)

var ratings = []int32{2, 4, 2, 6, 1, 7, 8, 9, 2, 1}

func candies(n int32, arr []int32) int64 {
	// Write your code here

	candies := make([]int64, n)

	// Create an array of length n and populate it with default int of 1
	for index := range candies {
		candies[index] = 1
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			candies[i] = candies[i]
			if candies[i+1]+1 > candies[i] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	var result int64

	for _, value := range candies {
		result += value
	}

	return result

}

func main() {
	result := candies(int32(len(ratings)), ratings)

	fmt.Println(result)
}
