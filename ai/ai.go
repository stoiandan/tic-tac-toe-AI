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

	symbol := getSymbol(b)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.GameBoard[i][j] == 0 {
				newState := translate(b, i, j, symbol)
				possibleFutureStates = append(possibleFutureStates, newState)
			}
		}
	}
	return possibleFutureStates
}

func getSymbol(b *model.Board) int {
	if b.IsPlayerX() == true {
		return 1
	}
	return -1
}

//translate  returnes a new baord, based on board b, where it applies the symbol on the board, at the given coordinates
func translate(b *model.Board, Xcoordinet, Ycoordinate int, symbol int) model.Board {
	newBoard := util.CloneBoard(b)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newBoard.GameBoard[i][j] = b.GameBoard[i][j]
		}
	}

	newBoard.GameBoard[Xcoordinet][Ycoordinate] = symbol

	return newBoard
}
