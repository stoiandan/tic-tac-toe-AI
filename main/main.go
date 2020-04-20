package main

import (
	"fmt"
	"tictactoe/ai"
	"tictactoe/model"
)

func main() {
	board := model.Board{}
	board = ai.MiniMax(&board)
	fmt.Println(board)

	_, score := board.IsGameOver()
	if score == 0 {
		fmt.Println("Score is even")
	}

	board.PreattyPrint()
}
