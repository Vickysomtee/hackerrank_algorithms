package main

import "fmt"

func staircase(n int) {
	strin := ""

	for i := 0; i < n; i++ {
		for j := 1; j < n+1; j++ {
			if j < n-i {
				strin += " "
			} else {
				strin += "#"
			}
		}
		fmt.Println(strin)
		strin = ""
	}
}

func main() {
	staircase(6)
}
