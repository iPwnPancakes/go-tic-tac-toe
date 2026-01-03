package tictactoe

import "strings"

type Rule interface {
	HasWon(board [][]string) bool
	GetDescription() string
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

func (r HorizontalMatch) GetDescription() string {
	return "HorizontalMatch"
}

type VerticalMatch struct{}

func (r VerticalMatch) HasWon(board [][]string) bool {
	var wholeStr string
	for row := 0; row < len(board); row++ {
		wholeStr += board[row][0]
	}

	return wholeStr == strings.Repeat(PlayerX, len(board)) || wholeStr == strings.Repeat(PlayerY, len(board))
}

func (r VerticalMatch) GetDescription() string {
	return "VerticalMatch"
}

type DiagonalMatch struct{}

func (r DiagonalMatch) GetDescription() string {
	return "DiagonalMatch"
}

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

type CornerMatch struct{}

func (r CornerMatch) GetDescription() string {
	return "CornerMatch"
}

func (r CornerMatch) HasWon(board [][]string) bool {
	length := len(board) - 1
	wholeStr := board[0][0] + board[0][length] + board[length][0] + board[length][length]

	return wholeStr == strings.Repeat(PlayerX, 4) || wholeStr == strings.Repeat(PlayerY, 4)
}
