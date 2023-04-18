package network

import (
	"tictactoe/game"
	"log"
)

type FlagPlayer = game.FlagPlayer
type Board = game.Board

type NetworkBoard struct {
	board *Board
	playerFlag map[*NetworkPlayer]FlagPlayer
}

func NewNetworkBoard() *NetworkBoard {
	return &NetworkBoard{
		board: game.NewBoard(),
		playerFlag: make(map[*NetworkPlayer]FlagPlayer),
	}
}

func (board *NetworkBoard) AddPlayer(player *NetworkPlayer) {
	log.Println(len(board.playerFlag))
	switch len(board.playerFlag) {
	case 0:
		board.playerFlag[player] = game.Player1
	case 1:
		board.playerFlag[player] = game.Player2
	}
}


func (board *NetworkBoard) RemovePlayer(player *NetworkPlayer) {
	if _, found := board.playerFlag[player]; !found {
		return
	}
	delete(board.playerFlag, player)
}

func (board *NetworkBoard) PlayTurn(player *NetworkPlayer, posx int, posy int) {
	board.board.Mark(board.playerFlag[player], posx, posy)
}

func (board *NetworkBoard) CheckWinner() *NetworkPlayer {
	winner_flag := board.board.CheckWinner()
	log.Println("Winner: ", winner_flag)
	for player, flag := range board.playerFlag {
		if flag == winner_flag {
			return player
		}
	}
	return nil
}
