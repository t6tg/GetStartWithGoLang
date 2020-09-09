package main

import "fmt"

// Global Variable
var count int = 0

func main() {
	score := 75
	switch {
	case score >= 75:
		fmt.Print("C+")
	case score >= 80:
		fmt.Print("B+")
	}
}
