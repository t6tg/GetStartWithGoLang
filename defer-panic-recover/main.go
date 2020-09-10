package main

import "fmt"

func main() {
	defer func() {
		v := recover() // print panic
		fmt.Print(v)
	}()
	defer a1()
	defer a2()                          // defer : work at end of function - stack ( connectDB , Redis , memcached )
	panic("Connect Database Error!!!!") // stop program : sensitive case , dangerous case
	a3()
}

func a1() {
	fmt.Println("A1")
}

func a2() {
	fmt.Println("A2")
}
func a3() {
	fmt.Println("A3")
}
