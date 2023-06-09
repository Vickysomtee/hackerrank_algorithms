package main

import "fmt"

var alice = []int{17, 28, 30}
var bob = []int{99, 16, 8}

// Second test case
// var alice = []int{6, 7, 8}
// var bob = []int{9, 7, 4}

func compareTriplet(a []int, b []int) []int {
	alice := 0
	bob := 0

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alice += 1
		} else if a[i] < b[i] {
			bob += 1
		}
	}

	return []int{alice, bob}

}

func main() {
	result := compareTriplet(alice, bob)

	fmt.Println(result)
}
