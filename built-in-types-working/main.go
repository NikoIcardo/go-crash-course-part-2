package main

import "fmt"

func main() {
	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)

	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call x is ", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
