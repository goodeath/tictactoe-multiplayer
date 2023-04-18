package game

type FlagPlayer int

const (
	Player1 FlagPlayer = -1
	Player2 FlagPlayer = 1
	NoWin FlagPlayer = 0
)

type Board struct {
	points [3][3]FlagPlayer
	player1 FlagPlayer
	player2 FlagPlayer
} 

func NewBoard() *Board {
	return &Board{
		player1: Player1,
		player2: Player2,
		points: [3][3]FlagPlayer{},

	}
}

func (b *Board) CheckWinner() FlagPlayer {
	winner := NoWin
	for i := 0; i < 3 && winner == NoWin; i++ {
		for j := 0; j < 3 && winner == NoWin; j++ {
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

func (board *Board) checkColumn(column int) FlagPlayer {
	win_one, win_two :=  true, true
	for i := 0; i < 3; i++ {
		colValue := board.points[column][i]
		if colValue == Player1 { 
			win_two = false
		} else if colValue == Player2 { 
			win_one = false 
		} else { 
			win_one, win_two = false, false
		}
	}

	if win_two { 
		return Player2
	} else if win_one {
		return Player1
	}
	return NoWin
}


func (board *Board) checkRow(row int) FlagPlayer {
	win_one, win_two :=  true, true
	for i := 0; i < 3; i++ {
		colValue := board.points[i][row]
		if colValue == Player1 { 
			win_two = false
		} else if colValue == Player2 { 
			win_one = false 
		} else { 
			win_one, win_two = false, false
		}
	}

	if win_two { 
		return Player2
	} else if win_one {
		return Player1
	}
	return NoWin
}


func (board *Board) checkDiagonal(row int) FlagPlayer {
	win_one, win_two :=  true, true
	for i := 0; i < 3; i++ {
		colValue := board.points[i][i]
		if colValue == Player1 { 
			win_two = false
		} else if colValue == Player2 { 
			win_one = false 
		} else { 
			win_one, win_two = false, false
		}
	}

	if win_two { 
		return Player2
	} else if win_one {
		return Player1
	}


	win_one, win_two =  true, true
	for i := 0; i < 3; i++ {
		colValue := board.points[i][2-i]
		if colValue == Player1 { 
			win_two = false
		} else if colValue == Player2 { 
			win_one = false 
		} else { 
			win_one, win_two = false, false
		}
	}

	return NoWin
}

func (b *Board) Mark(player FlagPlayer, i int, j int) {
	b.points[i][j] = player
}

func (b *Board) GetMarkAt(i int, j int) FlagPlayer {
	return b.points[i][j]
}
