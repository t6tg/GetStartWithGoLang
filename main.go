package main

import "fmt"

var m = 10

func main() {
	fmt.Println(sum(1, 2))
}

func sum(x int, y int) int {
	// m := 30
	fmt.Print(m)
	return x + y
}
