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

//CurrentPlayerSymbol returns the symbol of the current player - for X and -1 for 0
func CurrentPlayerSymbol(b *model.Board) model.Symbol {
	if b.IsPlayerX() == true {
		return model.X
	}
	return model.Zero
}
