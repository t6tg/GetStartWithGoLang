package main

import "fmt"

func main() {
	x := 10
	y := 20
	z := sum(x, y)
	// _, listNum := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	avg := av2(1, 2, 3, 4, 5)
	fmt.Println(z)
	fmt.Println(avg)
}

func sum(x, y int) int {
	return x + y
}

func average(xs []float64) (res float64) {
	var sum float64
	for _, v := range xs {
		sum = sum + v
	}
	res = sum / float64(len(xs))
	return res
}

func av2(xs ...float64) (res float64) {
	var sum float64
	for _, v := range xs {
		sum = sum + v
	}
	res = sum / float64(len(xs))
	return res
}
