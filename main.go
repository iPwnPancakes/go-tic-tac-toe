package main

import (
	"fmt"

	"github.com/ipwnpancakes/go-tic-tac-toe/tictactoe"
)

func main() {
	game := tictactoe.Game{}
	game.AddRule(tictactoe.HorizontalMatch{})
	game.AddRule(tictactoe.VerticalMatch{})
	game.AddRule(tictactoe.DiagonalMatch{})

	game.Start(3)

	for !game.HasWon() {
		fmt.Printf("It's %s's Turn.", game.GetCurrentPlayer())
		fmt.Println("Board State:")
		printBoard(game.GetBoard())

		var row, col int

		fmt.Println("Which row/col?")
		fmt.Scan(&row, &col)

		err := game.SetPiece(row, col)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !game.HasWon() {
			game.PassTurn()
		}
	}

	fmt.Printf("%s HAS WON WOOOOO", game.GetCurrentPlayer())
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
}
