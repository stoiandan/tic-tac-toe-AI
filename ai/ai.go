package ai

import (
	"tictactoe/model"
)

//Actions takes a state board and returns all possible states that can be transitioned to, from the input state
func Actions(b *model.Board) []model.Board {
	if isOver, _ := b.IsGameOver(); isOver == true {
		return nil
	}

	possibleFutureStates := make([]model.Board, 0, 9)

	symbol := b.Player()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.GameBoard[i][j] == 0 {
				newState := MakeMove(*b, model.Move{RowIndex: i, ColumnIndex: j, Symbol: symbol})
				possibleFutureStates = append(possibleFutureStates, newState)
			}
		}
	}
	return possibleFutureStates
}

//MakeMove  returnes a new baord, with the state of the old one, after applying a move to it.
func MakeMove(b model.Board, move model.Move) model.Board {
	b.GameBoard[move.RowIndex][move.ColumnIndex] = int(move.Symbol)

	return b
}

//MiniMax takes a board as input and returns a new board with the optimal move applied
func MiniMax(b *model.Board) model.Board {
	if isOver, _ := b.IsGameOver(); isOver == true {
		return *b
	}
	currentPlayer := b.Player()

	if currentPlayer == model.X {
		bestBoard := max(b)
		return MiniMax(&bestBoard)
	}
	bestBoard := min(b)
	return MiniMax(&bestBoard)
}

func min(b *model.Board) model.Board {

	bestCase := 2
	var bestBoard *model.Board

	for _, possibleBoard := range Actions(b) {
		result := MiniMax(&possibleBoard)
		if _, score := result.IsGameOver(); score < bestCase {
			bestCase = score
			bestBoard = &possibleBoard
		}
	}
	return *bestBoard
}

func max(b *model.Board) model.Board {

	bestCase := -3
	var bestBoard *model.Board

	for _, possibleBoard := range Actions(b) {
		result := MiniMax(&possibleBoard)
		if _, score := result.IsGameOver(); score > bestCase {
			bestCase = score
			bestBoard = &possibleBoard
		}
	}
	return *bestBoard
}
