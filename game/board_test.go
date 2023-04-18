package game

import (
	"testing"
)

func TestCheckMark(test *testing.T) {
	board := NewBoard()
	board.Mark(Player1, 0, 0)
	got := board.GetMarkAt(0, 0)
	expected := Player1
	if expected != got {
		test.Errorf("Expected: %v Received: %v", expected, got)
	}
}

func TestCheckWin(test *testing.T) {
	board := NewBoard()
	board.Mark(Player1, 0, 0)
	board.Mark(Player1, 1, 1)
	board.Mark(Player1, 2, 2)
	winner := board.CheckWinner()
	if winner != Player1 {
		test.Errorf("Expected: %v Received: %v", Player1, winner)
	}

	board = NewBoard()
	board.Mark(Player2, 0, 0)
	board.Mark(Player2, 1, 1)
	board.Mark(Player2, 2, 2)
	winner = board.CheckWinner()
	if winner != Player2 {
		test.Errorf("Expected: %v Received: %v", Player2, winner)
	}

	board = NewBoard()
	board.Mark(Player2, 0, 1)
	board.Mark(Player2, 1, 1)
	board.Mark(Player2, 2, 1)
	board.Mark(Player1, 0, 0)
	board.Mark(Player1, 1, 0)

	winner = board.CheckWinner()
	if winner != Player2 {
		test.Errorf("Expected: %v Received: %v", Player2, winner)
	}
}
