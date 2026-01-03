package main

import (
	"fmt"

	"github.com/ipwnpancakes/go-tic-tac-toe/tictactoe"
)

func main() {
	fmt.Println("Welcome to my implementation of tic-tac-toe in golang!")
	fmt.Println()

	game := tictactoe.Game{}
	game.AddRule(tictactoe.HorizontalMatch{})
	game.AddRule(tictactoe.VerticalMatch{})
	game.AddRule(tictactoe.DiagonalMatch{})
	game.AddRule(tictactoe.CornerMatch{})

	printRules(game.GetRules())

	err := game.Start(3, tictactoe.PlayerX)
	if err != nil {
		panic(err)
	}

	for !game.HasWon() {
		fmt.Printf("It's %s's Turn.\n\n", game.GetCurrentPlayer())
		fmt.Println("Board State:")
		printBoard(game.GetBoard())

		row, col := collectInput()

		err := game.SetPiece(row, col)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !game.HasWon() {
			game.PassTurn()
		}
	}

	fmt.Printf("%s HAS WON WOOOOO\n", game.GetCurrentPlayer())
}

func printRules(rules []tictactoe.Rule) {
	fmt.Println("Rules in play:")

	for _, rule := range rules {
		fmt.Println("- " + rule.GetDescription())
	}

	fmt.Println()
}

func printBoard(board [][]string) {
	for i, row := range board {
		for j, cell := range row {
			if cell == tictactoe.EmptyCell {
				fmt.Print("   ")
			} else {
				fmt.Printf(" %s ", cell)
			}
			if j < len(row)-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < len(board)-1 {
			for j := 0; j < len(row); j++ {
				fmt.Print("---")
				if j < len(row)-1 {
					fmt.Print("+")
				}
			}
			fmt.Println()
		}
	}

	fmt.Println()
}

func collectInput() (int, int) {
	var row, col int

	for {
		fmt.Println("Which row/col? Format: \"X Y\"")
		_, err := fmt.Scan(&row, &col)
		if err != nil {
			fmt.Println("Invalid input. Please enter two numbers separated by a space.")
			// Clear the input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		fmt.Println()
		break
	}

	return row, col
}
