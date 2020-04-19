package main

import (
	"fmt"
	"tictactoe/ai"
	"tictactoe/model"
)

func main() {
	board := model.Board{}
	endGame := ai.MiniMax(&board)

	fmt.Println(endGame)
}
