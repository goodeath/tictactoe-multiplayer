package tests

import (
	"testing"
	"tictactoe/game"
)

func TestCheckMark(test *testing.T) {
	var board game.Board
	board.Mark(board.FlagPlayer1, 0, 0)
	got := board.GetMarkAt(0, 0)
	expected := board.FlagPlayer1
	if expected != got {
		test.Errorf("Expected: %v Received: %v", expected, got)
	}
}

func TestCheckWin(test *testing.T) {
	board := new(game.Board).Init()
	board.Mark(board.FlagPlayer1, 0, 0)
	board.Mark(board.FlagPlayer1, 1, 1)
	board.Mark(board.FlagPlayer1, 2, 2)
	winner := board.CheckWinner()
	if winner != board.FlagPlayer1 {
		test.Errorf("Expected: %v Received: %v", board.FlagPlayer1, winner)
	}

	board = new(game.Board).Init()
	board.Mark(board.FlagPlayer2, 0, 0)
	board.Mark(board.FlagPlayer2, 1, 1)
	board.Mark(board.FlagPlayer2, 2, 2)
	winner = board.CheckWinner()
	if winner != board.FlagPlayer2 {
		test.Errorf("Expected: %v Received: %v", board.FlagPlayer2, winner)
	}

	board = new(game.Board).Init()
	board.Mark(board.FlagPlayer2, 0, 1)
	board.Mark(board.FlagPlayer2, 1, 1)
	board.Mark(board.FlagPlayer2, 2, 1)
	board.Mark(board.FlagPlayer1, 0, 0)
	board.Mark(board.FlagPlayer1, 1, 0)

	winner = board.CheckWinner()
	if winner != board.FlagPlayer2 {
		test.Errorf("Expected: %v Received: %v", board.FlagPlayer2, winner)
	}
}
