package main

import (
	"fmt"
	"tictactoe/ai"
	"tictactoe/model"
)

func main() {
	board := model.Board{}
	isOver, _ := board.IsGameOver()
	for ; isOver == false; isOver, _ = board.IsGameOver() {
		fmt.Println(board)
		board = ai.MiniMax(&board)
	}
	fmt.Println(board)

	_, score := board.IsGameOver()
	fmt.Println("Winner is:", score)
}
