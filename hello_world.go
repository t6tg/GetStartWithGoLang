package main

import "fmt"

func main() {
	// Variable
	a := map[string]string{
		"Hello": "Gopher",
		"Name":  "James",
		"X":     "Hello from the other side.",
	}

	if x, ok := a["X"]; ok {
		fmt.Println(x)
	}
	for key, val := range a {
		fmt.Println(key, ":", val)
	}

}
