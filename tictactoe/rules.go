package tictactoe

type Rule interface {
	HasWon(board [][]string) bool
}

type HorizontalMatch struct{}

func (r HorizontalMatch) HasWon(board [][]string) bool {
	for row := 0; row < len(board); row++ {
		if board[row][0] != EmptyCell && allSameInRow(board[row]) {
			return true
		}
	}

	return false
}

func allSameInRow(arr []string) bool {
	first := arr[0]
	for _, s := range arr[1:] {
		if s != first {
			return false
		}
	}

	return true
}

type VerticalMatch struct{}

func (r VerticalMatch) HasWon(board [][]string) bool {
	if board[0][0] == EmptyCell {
		return false
	}

	var lastChar string
	for row := 0; row < len(board); row++ {
		if row == 0 {
			lastChar = board[row][0]

			continue
		}

		if board[row][0] != lastChar {
			return false
		}
	}

	return true
}
