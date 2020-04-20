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

	bestBoard := ai.MiniMax(&board)

	if bestBoard.GameBoard[2][2] == 0 {
		t.Errorf("Min does not make the best decission")
	}
}

func TestDoesMaxMakeBestMove(t *testing.T) {
	board := model.Board{GameBoard: [3][3]int{{1, 1, -1},
		{-1, 0, -1},
		{1, 0, 0}}}

	bestBoard := ai.MiniMax(&board)

	if _, score := bestBoard.IsGameOver(); score != 0 {
		t.Errorf("Max does not make the best decission")
	}
}

func TestDoesMinPreventMaxFromWinning(t *testing.T) {
	board := model.Board{GameBoard: [3][3]int{{1, -1, 1},
		{-1, 1, 0},
		{0, 0, 0}}}

	bestBoard := ai.MiniMax(&board)

	if bestBoard.GameBoard[0][2] == 0 {
		t.Errorf("Min is an idiot")
	}
}
