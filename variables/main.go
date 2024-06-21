package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when ready."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	printGame(firstNumber, secondNumber, subtraction, answer)

}

func printGame(firstNumber int, secondNumber int, subtraction int, answer int) {
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

	fmt.Println("The answser is", answer)
}
