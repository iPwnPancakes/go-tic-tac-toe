package tictactoe

type Rule interface {
	HasWon(board [][]string) bool
}

type HorizontalMatch struct{}

func (r HorizontalMatch) HasWon(board [][]string) bool {
	for row := 0; row < len(board); row++ {
		if board[row][0] != "" && allSameInRow(board[row]) {
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
