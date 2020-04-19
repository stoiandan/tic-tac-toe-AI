package tests

import (
	"testing"
	"tictactoe/ai"
	"tictactoe/model"
)

func TestDoesMinMakeBestMove(t *testing.T) {
	board := model.Board{GameBoard: [3][3]int{{1, 1, -1},
		{-1, 1, -1},
		{1, 0, 0}}}

	bestBoard := ai.Min(&board)

	if bestBoard.GameBoard[2][2] == 0 {
		t.Errorf("Min does not make the best decission")
	}
}
