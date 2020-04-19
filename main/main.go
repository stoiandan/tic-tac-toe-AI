package main

import (
	"fmt"
	"tictactoe/ai"
	"tictactoe/model"
)

func main() {
	board := model.Board{}
	isOver, _ := board.IsGameOver()
	for isOver == false {
		fmt.Println(board)
		board = ai.MiniMax(&board)
		isOver, _ = board.IsGameOver()
	}
}
