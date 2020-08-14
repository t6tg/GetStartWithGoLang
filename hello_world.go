package main

import (
	"fmt"
)

func main() {
	// Variable
	var gopher string
	gopher = "Gopher"
	fmt.Printf("Hello, %s\n", gopher)

	name := "Thanawat Gulati"
	fmt.Printf("My name is %s\n", name)
	// Decision
	var fruit string
	fmt.Print("Enter Fruit : ")
	fmt.Scan(&fruit)
	if len(fruit) > 0 {
		fmt.Printf("Fruit = %s ", fruit)
	}
	switch fruit {
	case "apple":
		fmt.Println("🍎")
	case "banana":
		fmt.Println("🍌")
	case "mango":
		fmt.Println("🥭")
	default:
		fmt.Println("💩")
	}

}
