package main

// type person struct {
// 	Name string
// 	age  int
// }

// func (p person) mutatePerson(name string) {
// 	p.Name = name
// 	fmt.Println("inside mutate:", p)
// }

func main() {
	// p1 := person{Name: "James"}
	// fmt.Println(p1)
	// p1.mutatePerson("Hacker")
	// fmt.Println(p1)
	// a := [5]int{}
	// mutateArray(a[:]) // slice is a part of pointer
	// fmt.Println(a)
	// fmt.Println(add(1, 3))
	// Variable
	// a := map[string]string{
	// 	"Hello": "Gopher",
	// 	"Name":  "James",
	// 	"X":     "Hello from the other side.",
	// }
	// if x, ok := a["X"]; ok {
	// 	fmt.Println(x)
	// }
	// for key, val := range a {
	// 	fmt.Println(key, ":", val)
	// }
	// Pointer
	// a := 10
	// fmt.Println(a)
	// ptrA := &a
	// fmt.Println(ptrA)
	// fmt.Println(*ptrA)
	// *ptrA = 15
	// fmt.Println(a)

}
func add(x, y int) int {
	p := 1
	return x + y + p
}

func mutateArray(a []int) {
	a[0] = 10
}
