package main

import "fmt"

func main() {
	var i int
	var p *int
	i = 10
	fmt.Printf("i = %d\n", i)
	fmt.Printf("Address i = %v\n", &i)
	p = &i
	fmt.Printf("Address i in P = %v\n", p)
	fmt.Printf("Address of P = %v\n", &p)
	fmt.Printf("value of P = %v\n", *p)
}
