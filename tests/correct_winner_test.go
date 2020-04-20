package tests

import (
	"testing"
	"tictactoe/model"
)

func TestCorretWinnerIsReported(t *testing.T) {
	board := model.Board{GameBoard: [3][3]int{{0, 0, 1},
		{-1, 1, -1},
		{1, -1, 1}}}

	if _, winner := board.IsGameOver(); winner == -1 {
		t.Errorf("False winner is reported")
	}
}
