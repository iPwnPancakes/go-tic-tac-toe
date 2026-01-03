package tictactoe

import "strings"

type Rule interface {
	HasWon(board [][]string) bool
}

type HorizontalMatch struct{}

func (r HorizontalMatch) HasWon(board [][]string) bool {
	for row := 0; row < len(board); row++ {
		wholeStr := ""

		for col := 0; col < len(board[row]); col++ {
			wholeStr += board[row][col]
		}

		if wholeStr == strings.Repeat(PlayerX, len(board)) || wholeStr == strings.Repeat(PlayerY, len(board)) {
			return true
		}
	}

	return false
}

type VerticalMatch struct{}

func (r VerticalMatch) HasWon(board [][]string) bool {
	var wholeStr string
	for row := 0; row < len(board); row++ {
		wholeStr += board[row][0]
	}

	return wholeStr == strings.Repeat(PlayerX, len(board)) || wholeStr == strings.Repeat(PlayerY, len(board))
}

type DiagonalMatch struct{}

func (r DiagonalMatch) HasWon(board [][]string) bool {
	var wholeStr string

	wholeStr = ""
	for row := 0; row < len(board); row++ {
		wholeStr += board[row][row]
	}

	if wholeStr == strings.Repeat(PlayerX, len(board)) || wholeStr == strings.Repeat(PlayerY, len(board)) {
		return true
	}

	wholeStr = ""
	for row := 0; row < len(board); row++ {
		wholeStr += board[row][len(board)-1-row]
	}

	return wholeStr == strings.Repeat(PlayerX, len(board)) || wholeStr == strings.Repeat(PlayerY, len(board))
}
