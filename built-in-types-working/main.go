package main

import "fmt"

type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (d *Cat) Says() string {
	return d.Sound
}

func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}

func main() {
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	riddle(&dog)

	var cat Cat

	cat.Name = "cat"

	cat.NumberOfLegs = 4
	cat.HasTail = true
	cat.Sound = "meow"

	riddle(&cat)
}

func riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
