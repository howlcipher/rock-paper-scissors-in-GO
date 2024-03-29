package main

import (
	"fmt"
	"math/rand"
	"time"
)

//choices in game
var choices = [3]string{"Rock", "Paper", "Scissors"}

//main
func main() {
	fmt.Println("Welcome to Rock Paper Scissors")
	var userChoice string = choice()           //choice of the user
	var computerChoice string = randomChoice() // choice of the computer
	pickWinner(userChoice, computerChoice)     // winnder is chosen
}

//list all choice in cli
func listChoices() {
	message := fmt.Sprintf("Please choose \n1. %s\n2. %s\n3. %s", choices[0], choices[1], choices[2])
	fmt.Println(message)
}

// get the choice by user or computer
func choice() string {
	var choice int
	var chosen string
	for i := true; i != false; {
		listChoices()
		fmt.Scan(&choice)

		if choice == 1 {
			chosen = choices[choice-1]
			printChoice(chosen, "user")
			i = false
		} else if choice == 2 {
			chosen = choices[choice-1]
			printChoice(chosen, "user")
			i = false
		} else if choice == 3 {
			chosen = choices[choice-1]
			printChoice(chosen, "user")
			i = false
		} else {
			fmt.Println("Invalid choice try again...choose 1, 2,or 3.")
			continue
		}

	}
	return chosen
}

// print choice select
func printChoice(choice string, user string) {
	fmt.Println(user, "choice is...", choice)
}

//random choice
func randomChoice() string {
	rand.Seed(time.Now().UnixNano())
	var random int = rand.Intn(3)
	printChoice(choices[random], "Computer")
	return choices[random]
}

// decide who the winner is
func pickWinner(userChoice string, computerChoice string) {

	if userChoice == "Rock" && computerChoice == "Scissors" {
		fmt.Println("User wins!")
	} else if userChoice == "Paper" && computerChoice == "Rock" {
		fmt.Println("User wins!")
	} else if userChoice == "Scissors" && computerChoice == "Paper" {
		fmt.Println("User wins!")
	} else if userChoice == computerChoice {
		fmt.Println("Draw!")
	} else {
		fmt.Println("Computer wins!")
	}

}
