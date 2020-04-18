package ai

import (
	"tictactoe/model"
	"tictactoe/util"
)

//Actions takes a state board and returns all possible states that can be transitioned to, from the input state
func Actions(b *model.Board) []model.Board {
	if isOver, _ := b.IsGameOver(); isOver == true {
		return nil
	}

	possibleFutureStates := make([]model.Board, 0, 9)

	symbol := util.CurrentPlayerSymbol(b)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.GameBoard[i][j] == 0 {
				newState := MakeMove(b, model.Move{RowIndex: i, ColumnIndex: j, Symbol: symbol})
				possibleFutureStates = append(possibleFutureStates, newState)
			}
		}
	}
	return possibleFutureStates
}

//MakeMove  returnes a new baord, with the state of the old one, after applying a move to it.
func MakeMove(b *model.Board, move model.Move) model.Board {
	newBoard := util.CloneBoard(b)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newBoard.GameBoard[i][j] = b.GameBoard[i][j]
		}
	}

	newBoard.GameBoard[move.RowIndex][move.ColumnIndex] = int(move.Symbol)

	return newBoard
}
