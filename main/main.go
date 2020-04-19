package main

import (
	"tictactoe/ai"
	"tictactoe/model"
)

func main() {
	board := model.Board{}
	ai.MiniMax(&board)
}
