package main

import "fmt"

func main() {
	// Variable
	a := make(map[string]string)
	a["Hello"] = "Gopher"
	a["Name"] = "James"
	a["X"] = "Hello from the other side."

	if x, ok := a["X"]; ok {
		fmt.Println(x)
	}

}
