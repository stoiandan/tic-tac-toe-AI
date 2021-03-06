package model

import "fmt"

// Board represents the game GameBoard
type Board struct {
	GameBoard [3][3]int
}

//Player returns the Symbol of who's turn is it(X or 0)
func (b Board) Player() Symbol {
	/*
		X is represented as 1
		0 is represented as -1
		empty cell is 0
	*/
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += b.GameBoard[i][j]
		}
	}

	// This is the beginning of the game, presupose X starts
	// or there are more 0 (-1's) on the GameBoard
	if sum <= 0 {
		return X
	}
	return Zero
}

//IsGameOver reutnrs a tuple (bool,int) telling if the game is over, and if it is, who won
func (b Board) IsGameOver() (bool, int) {

	isAnyFieldEmpty := false
	i, j := 0, 0
	for i = 0; i < 3; i++ {
		rowScore := 0
		columnScore := 0
		for j = 0; j < 3; j++ {
			if b.GameBoard[i][j] == 0 {
				isAnyFieldEmpty = true
			}
			rowScore += b.GameBoard[i][j]
			columnScore += b.GameBoard[j][i]
		}
		if rowScore == 3 || rowScore == -3 {
			return true, rowScore / 3
		}
		if columnScore == 3 || columnScore == -3 {
			return true, columnScore / 3
		}
	}

	// check the main diagonal
	if b.GameBoard[0][0] == b.GameBoard[1][1] && b.GameBoard[1][1] == b.GameBoard[2][2] && b.GameBoard[1][1] != 0 {
		return true, b.GameBoard[1][1]
	}

	//check the secodnary diagonal
	if b.GameBoard[0][2] == b.GameBoard[1][1] && b.GameBoard[1][1] == b.GameBoard[2][0] && b.GameBoard[1][1] != 0 {
		return true, b.GameBoard[1][1]
	}

	// game is still going on
	if isAnyFieldEmpty == true {
		return false, 0
	}

	// is score even
	return true, 0
}

//PreattyPrint pritns the GamebBorad in terms of X and 0 as opposed to mathematical model
func (b Board) PreattyPrint() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch b.GameBoard[i][j] {
			case -1:
				fmt.Print("0 ")
			case 1:
				fmt.Print("X ")
			case 0:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
