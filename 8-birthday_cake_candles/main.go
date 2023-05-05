package main

import "fmt"

// Test case 1
var candles = []int{3, 2, 1, 3}

// Test case 2
// var candles = []int{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}

// This function gets the max value in a slice
// Go Lang doesn't have a built in function for get the max value
func getMax(arr []int) int {
	number := 0
	for _, value := range arr {
		if value > number {
			number = value
		}
	}

	return number
}

func birthdayCakeCandles(candles []int) int {
	max := getMax(candles)
	count := 0

	for _, value := range candles {
		if value == max {
			count++
		}
	}

	return count
}

func main() {
	result := birthdayCakeCandles(candles)

	fmt.Println(result)
}
