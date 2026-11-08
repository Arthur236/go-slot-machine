package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Println("Welcome to The Casino!")
	fmt.Printf("Enter your name:")

	_, err := fmt.Scanln(&name)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance = %d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("You cannot bet more than your balance.")
		} else {
			break
		}
	}

	return bet
}
