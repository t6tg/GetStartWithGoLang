package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start .....")
	// log.Fatal("Error ......") // Stop program
	doSaveWork()
	fmt.Println("Done")

}

func doSaveWork() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Fail : ", r)
		}
	}()
	doFailWork()
	fmt.Println("Work success")
}

func doFailWork() {
	panic("fail")
}
