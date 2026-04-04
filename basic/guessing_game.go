package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	radom := rand.New(source)

	// Generate a random number between 1 and 100
	target := radom.Intn(100) + 1

	// welcome message
	fmt.Println("Welcome to the Guessing Game")
	fmt.Println("I have a number between 1 and 100")
	fmt.Println("Can you guess what is it?")

	var guess int
	// 伪while go里面没有while 循环 这里for loop 实现类似其他语言的while
	for {
		fmt.Println("Please enter your guess:")
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Println("Congratulations, you guessed the correct number!")
			break
		} else if guess > target {
			fmt.Println("Too high! Try guessing a lower number.")
		} else {
			fmt.Println("Too low! Try guessing a higher number.")
		}
	}

}
