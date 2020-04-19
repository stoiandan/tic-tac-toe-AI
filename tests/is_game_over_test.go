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
