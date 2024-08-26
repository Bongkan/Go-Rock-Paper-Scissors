package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var choices = []string{"rock", "paper", "scissors"}
	rand.Seed(time.Now().UnixNano())
	// choice := choices[rand.Intn(len(choices))]
	userScore := 0
	computerScore := 0
	println("welcome to rock, paper, scissors")

	for {
		userChoices := ""
		println("enter your choice: rock, paper, or scissors (type exit to quit)")
		fmt.Scanln(&userChoices)
		isValidChoice := false
		for _, choice := range choices {
			if userChoices == choice {
				isValidChoice = true
				break
			}
			if !isValidChoice && userChoices != "exit" {
				println("invalid choice, please try again rock paper scissors only!")
				continue
			}
		}
		computerChoice := choices[rand.Intn(len(choices))]
		if userChoices == computerChoice {
			println("it's a tie!")
		}

		if userChoices == "rock" && computerChoice == "scissors" {
			println("you win!")
			userScore++
		} else if userChoices == "paper" && computerChoice == "rock" {
			println("you win!")
			userScore++
		} else if userChoices == "scissors" && computerChoice == "paper" {
			println("you win!")
			userScore++
		} else if userChoices == "exit" {
			println("goodbye!")
			break
		} else {
			println("you lose!")
			computerScore++
		}
		println("your score:", userScore)
		println("computer score:", computerScore)
		if userScore == 3 {
			println("you win the game!")
			break
		} else if computerScore == 3 {
			println("you lose the game!")
			break
		}
	}
}
