package util

import "tictactoe/model"

//CloneBoard createsa new board, that replicates the given one
func CloneBoard(b *model.Board) model.Board {
	newBoard := model.Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.GameBoard[i][j] = newBoard.GameBoard[i][j]
		}
	}
	return newBoard
}
