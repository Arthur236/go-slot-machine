package main

import (
	"fmt"
)

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	winningLines := []uint{}

	// Check for winning combinations (simplified example - check rows)
	for _, row := range spin {
		win := true
		checkSymbol := row[0]

		// Check if all symbols in the row are the same starting from the second symbol
		for _, symbol := range row[1:] {
			if symbol != checkSymbol {
				win = false
				break
			}
		}

		if win {
			winningLines = append(winningLines, multipliers[checkSymbol])
		} else {
			winningLines = append(winningLines, 0)
		}
	}

	return winningLines
}

func main() {
	symbols := map[string]uint{
		"♠": 4,
		"♥": 7,
		"♦": 12,
		"♣": 20,
	}

	// Multipliers for each symbol line
	multipliers := map[string]uint{
		"♠": 20,
		"♥": 10,
		"♦": 5,
		"♣": 2,
	}

	balance := uint(200)

	GetName()

	symbolArray := GenerateSymbolArray(symbols)

	for balance > 0 {
		bet := GetBet(balance)

		if bet == 0 {
			fmt.Println("Thanks for playing!")
			break
		}

		balance -= bet

		spin := GetSpin(symbolArray, 3, 3)

		PrintSpin(spin)

		// Check for winning combinations and update balance accordingly
		winningLines := checkWin(spin, multipliers)

		for i, multiplier := range winningLines {
			winAmount := bet * multiplier
			balance += winAmount

			if multiplier > 0 {
				fmt.Printf("You won %d, (%dx) on Line #%d!\n", winAmount, multiplier, i+1)
			}
		}

		fmt.Println("Current balance:", balance)
	}

	fmt.Println("You left with", balance)
}
