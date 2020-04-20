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

	symbol := b.Player()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.GameBoard[i][j] == 0 {
				newState := util.MakeMove(*b, model.Move{RowIndex: i, ColumnIndex: j, Symbol: symbol})
				possibleFutureStates = append(possibleFutureStates, newState)
			}
		}
	}
	return possibleFutureStates
}

//MiniMax takes a board as input and returns a new board with the optimal move applied
func MiniMax(b *model.Board) model.Board {
	if isOver, _ := b.IsGameOver(); isOver == true {
		return *b
	}
	currentPlayer := b.Player()

	if currentPlayer == model.X {
		result := max(b)
		return MiniMax(&result)
	}
	result := min(b)
	return MiniMax(&result)
}

func min(b *model.Board) model.Board {
	bestScore := 3
	var bestBoard model.Board

	for _, possibleBoard := range Actions(b) {
		board := MiniMax(&possibleBoard)
		if _, score := board.IsGameOver(); score < bestScore {
			bestScore = score
			bestBoard = possibleBoard
		}
	}
	return bestBoard
}

func max(b *model.Board) model.Board {
	bestScore := -3
	var bestBoard model.Board

	for _, possibleBoard := range Actions(b) {
		board := MiniMax(&possibleBoard)
		if _, score := board.IsGameOver(); score > bestScore {
			bestScore = score
			bestBoard = possibleBoard
		}
	}
	return bestBoard
}
