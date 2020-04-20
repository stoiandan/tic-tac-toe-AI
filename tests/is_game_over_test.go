package tests

import (
	"testing"
	"tictactoe/model"
)

func TestEmptyBoardIsGameOver(t *testing.T) {
	emptyBoard := model.Board{}

	if isOver, _ := emptyBoard.IsGameOver(); isOver == true {
		t.Error("Empty board returns gamve over")
	}
}

//If this test fails, the diagonals are not checked beore checking if the game is not yet over
// thus yielding the game is not over, when in fact it is
func TestBoardNotFullButGameNotOver(t *testing.T) {
	board := model.Board{GameBoard: [3][3]int{{1, 1, -1},
		{-1, 1, -1},
		{1, 0, 1}}}

	if isOver, _ := board.IsGameOver(); isOver == false {
		t.Errorf("Game is over false negative")
	}
}

func TestSecondDiagonalForGameOver(t *testing.T) {
	board := model.Board{GameBoard: [3][3]int{{0, 0, 1},
		{-1, 1, -1},
		{1, -1, 1}}}

	if isOver, _ := board.IsGameOver(); isOver == false {
		t.Errorf("Game over false negative of second diagonal")
	}
}
