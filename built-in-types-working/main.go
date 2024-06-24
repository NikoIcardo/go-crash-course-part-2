package main

import "log"

var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

func main() {
	myInt = 10
	myUint = 20
	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	// strings are immutable
	myString := "Trevor"

	log.Println(myString)

	// this creates a brand new string
	myString = "John"

	var myBool = true

	log.Println(myBool)
}
