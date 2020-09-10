package main

import "fmt"

func main() {
	score := [][]int{{20, 30}, {30, 40}}
	for key, val := range score {
		fmt.Println("Val :", val[1])
		fmt.Println("Key :", key)
	}
}
