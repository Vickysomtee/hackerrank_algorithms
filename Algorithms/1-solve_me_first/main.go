package main

import (
	"fmt"
)

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	result := solveMeFirst(2, 3)

	fmt.Println(result)
}
