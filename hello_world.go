package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("Result = %d\n", add(1, 3))
}
