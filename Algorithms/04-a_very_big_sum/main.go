package main

import "fmt"

var numbers = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

func aVeryBigSum(ar []int64) int64 {
	var result int64 = 0

	for i := range ar {
		result += ar[i]
	}

	return result
}

func main() {

	result := aVeryBigSum(numbers)

	fmt.Println(result)

}
