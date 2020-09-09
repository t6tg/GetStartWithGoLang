package main

import "fmt"

// Global Variable
var count int = 0

func main() {
	// Variable in Golang
	fmt.Println("Hello World.")
	var str string = "James"
	var boolean bool = true
	var Integer int = 1
	// implecit declration
	value := "James" // <- Dynamic Type
	run()
	fmt.Printf("String: %s ,Boolean: %t ,Integet: %d ,Implecit var: %v ,Global: %d \n", str, boolean, Integer, value, count)
}

func run() {
	// get Global variable
	count++
}
