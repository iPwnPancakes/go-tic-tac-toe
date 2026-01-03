package tictactoe

import "errors"

type Game struct {
	Size  int
	board [][]string
	turn  string
	rules []Rule
}

func (g *Game) Start(size int) {
	g.board = make([][]string, size)
	for row := 0; row < size; row++ {
		g.board[row] = make([]string, size)

		for col := 0; col < size; col++ {
			g.board[row][col] = " "
		}
	}

	g.turn = "X"
}

func (g *Game) AddRule(rule Rule) {
	g.rules = append(g.rules, rule)
}

func (g Game) GetBoard() [][]string {
	return g.board
}

func (g Game) GetCurrentPlayer() string {
	return g.turn
}

func (g *Game) PassTurn() {
	if g.turn == "X" {
		g.turn = "O"
	} else {
		g.turn = "X"
	}
}

func (g *Game) SetPiece(row int, col int) error {
	rows := len(g.board)
	if row >= rows {
		return errors.New("Row too big")
	}

	cols := len(g.board[0])
	if col >= cols {
		return errors.New("Col too big")
	}

	g.board[row][col] = g.turn

	return nil
}

func (g Game) HasWon() bool {
	for _, rule := range g.rules {
		if rule.HasWon(g.board) {
			return true
		}
	}

	return false
}
