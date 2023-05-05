package main

import "fmt"

var numbers = []int{2, 3, 4}

func simpleArraySum(arr []int) int {
	var sum int

	for _, value := range arr {
		sum += value
	}

	// Solution 2
	// for i := 0; i < len(arr); i++ {
	// 	sum += arr[i]
	// }

	return sum
}

func main() {
	result := simpleArraySum([]int{3, 4})

	fmt.Println(result)
}

// What I learnt ?

// ==> Short-hand variable syntax (:=) cannot be used outside a function scope.
