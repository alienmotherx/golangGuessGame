package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Introducing the user to the game
	var name string
	fmt.Println("Welcome to Guess The Number")
	fmt.Printf("Please input your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name, ("\n This is the number guessing game.\n You have 3 chances to guess the right number."))

	//Generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	//The Game
	var guess int
	n := 3
	for n >= 1 {
		fmt.Println("Please input your guess: ")
		fmt.Scan(&guess)
		if guess > secretNumber && n > 1 {
			fmt.Println("Too Big")
			n = n - 1
			fmt.Println("Chances Left:", n)
		} else if guess < secretNumber && n > 1 {
			fmt.Println("Too Small")
			n = n - 1
			fmt.Println("Chances Left:", n)
		} else if guess == secretNumber {
			fmt.Println("You WON!!! :) :)")
			break
		} else {
			fmt.Println("Chances Left:", 0)
			fmt.Println("Oh No :( :( . Wrong guesses, try again.")
			break
		}
	}

}
