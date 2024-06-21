package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	// var answer int

	reader := bufio.NewReader(os.Stdin)
	// display a welcom/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Think of a number between one and ten", prompt)

	reader.ReadString('\n')

	// take them through the game

	fmt.Println("Multiply your number by", firstNumber, prompt)

	fmt.Println("Now multiply the result by", secondNumber, prompt)

	fmt.Println("Divide the result by the number you originally thought of", prompt)

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
}
