package main

import "fmt"

func main() {
	// Variable in Golang
	fmt.Println("Hello World.")
	var str string = "James"
	var boolean bool = true
	var Integer int = 1
	// implecit declration
	value := "James" // <- No Type
	fmt.Printf("String: %s ,Boolean: %t ,Integet: %d , Implecit var: %v \n", str, boolean, Integer, value)
}
