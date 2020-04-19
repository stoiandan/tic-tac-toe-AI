package model

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

//IsGameOver reutnrs a tuple who's two values signify: the first values is a boolean, true if the game is over, false otherwise.
// The second value is an integer,   -1 if 0 won, 1 is X won, or 0 is the score is equal, no one won, -2 if the game is not over
func (b Board) IsGameOver() (bool, int) {

	isAnyFieldEmpty := false
	i, j := 0, 0
	for ; i < 3; i++ {
		rowScore := 0
		columnScore := 0
		for ; j < 3; j++ {
			if b.GameBoard[i][j] == 0 {
				isAnyFieldEmpty = true
			}
			rowScore += b.GameBoard[i][j]
			columnScore += b.GameBoard[j][i]
		}
		if rowScore == 3 || rowScore == -3 {
			return true, b.GameBoard[i][j]
		}
		if columnScore == 3 || columnScore == -3 {
			return true, b.GameBoard[j][i]
		}
	}
	// check the main diagonal
	if b.GameBoard[0][0] == b.GameBoard[1][1] && b.GameBoard[1][1] == b.GameBoard[2][2] {
		return true, b.GameBoard[1][1]
	}

	//check the secodnary diagonal
	if b.GameBoard[0][2] == b.GameBoard[1][2] && b.GameBoard[1][2] == b.GameBoard[2][0] {
		return true, b.GameBoard[1][2]
	}

	// is score even
	if isAnyFieldEmpty == false {
		return true, 0
	}
	// game is still going on
	return false, -2
}
