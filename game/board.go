package game

type Board struct {
	points [3][3]int
	FlagPlayer1 int
	FlagPlayer2 int
} 

func (b Board) Init() Board {
	b.FlagPlayer1 = -1
	b.FlagPlayer2 = 1
	return b
}

func (b *Board) CheckWinner() int {
	winner := 0
	for i := 0; i < 3 && winner==0; i++ {
		for j := 0; j < 3 && winner==0; j++ {
			switch {
			case i == 0:
				winner = b.checkColumn(j)
			case i == 1:
				winner = b.checkRow(j)
			case i == 2:
				winner = b.checkDiagonal(j)
			}
		}
	}
	return winner
}

func (board *Board) checkColumn(column int) int {
	win_one, win_two :=  true, true
	for i := 0; i < 3; i++ {
		colValue := board.points[i][column]
		if colValue == board.FlagPlayer1 { 
			win_two = false
		} else if colValue == board.FlagPlayer2 { 
			win_one = false 
		} else { 
			win_one, win_two = false, false
		}
	}

	if win_two { 
		return board.FlagPlayer2
	} else if win_one {
		return board.FlagPlayer1
	}
	return 0
}


func (board *Board) checkRow(row int) int {
	win_one, win_two :=  true, true
	for i := 0; i < 3; i++ {
		colValue := board.points[row][i]
		if colValue == board.FlagPlayer1 { 
			win_two = false
		} else if colValue == board.FlagPlayer2 { 
			win_one = false 
		} else { 
			win_one, win_two = false, false
		}
	}

	if win_two { 
		return board.FlagPlayer2
	} else if win_one {
		return board.FlagPlayer1
	}
	return 0
}


func (board *Board) checkDiagonal(row int) int {
	win_one, win_two :=  true, true
	for i := 0; i < 3; i++ {
		colValue := board.points[i][i]
		if colValue == board.FlagPlayer1 { 
			win_two = false
		} else if colValue == board.FlagPlayer2 { 
			win_one = false 
		} else { 
			win_one, win_two = false, false
		}
	}

	if win_two { 
		return board.FlagPlayer2
	} else if win_one {
		return board.FlagPlayer1
	}


	win_one, win_two =  true, true
	for i := 0; i < 3; i++ {
		colValue := board.points[i][2-i]
		if colValue == board.FlagPlayer1 { 
			win_two = false
		} else if colValue == board.FlagPlayer2 { 
			win_one = false 
		} else { 
			win_one, win_two = false, false
		}
	}

	return 0
}

func (b *Board) Mark(player int, i int, j int) {
	b.points[i][j] = player
}

func (b *Board) GetMarkAt(i int, j int) int {
	return b.points[i][j]
}
