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

//MiniMax returns the optimal move the player can make on a current state
func MiniMax(b *model.Board) model.Board {
	// current player
	player := b.Player()

	if player == model.X {
		return Max(b)
	}
	return Min(b)
}

//Min applies the best move that can be achieved from the current state, and return the board yilded from applying that mvoe
func Min(b *model.Board) model.Board {
	if IsOver, _ := b.IsGameOver(); IsOver == true {
		return *b
	}

	// We assume the best decisison is the worst
	// by starting at the worst we are guaranteed for find something better or at least the same
	bestDecission := 1
	var bestPossibleBoard *model.Board
	for _, possibleBoard := range Actions(b) {
		board := Min(&possibleBoard)
		if isOver, score := board.IsGameOver(); isOver == true && score < bestDecission {
			bestPossibleBoard = &board
		}
	}

	return *bestPossibleBoard
}

//Max applies the best move that can be achieved from the current state, and return the board yilded from applying that mvoe
func Max(b *model.Board) model.Board {
	if IsOver, _ := b.IsGameOver(); IsOver == true {
		return *b
	}

	// We assume the best decisison is the worst
	// by starting at the worst we are guaranteed for find something better or at least the same
	bestDecission := -1
	var bestPossibleBoard *model.Board
	for _, possibleBoard := range Actions(b) {
		board := Min(&possibleBoard)
		if isOver, score := board.IsGameOver(); isOver == true && score > bestDecission {
			bestPossibleBoard = &board
		}
	}

	return *bestPossibleBoard
}
