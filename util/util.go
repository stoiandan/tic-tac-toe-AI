package util

import (
	"tictactoe/model"
)

//MakeMove  returnes a new board, applying a move on top of the old board
func MakeMove(b model.Board, move model.Move) model.Board {
	b.GameBoard[move.RowIndex][move.ColumnIndex] = int(move.Symbol)
	return b
}
