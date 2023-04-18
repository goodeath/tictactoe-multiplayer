// Handle game through network
package network

import (
	"sync"
	"golang.org/x/exp/map"
)

type Game struct {
	playersInGame int
	players map[*NetworkPlayer]bool
	board *NetworkBoard
	turn *NetworkPlayer
	winner *NetworkPlayer
	sync.RWMutex
}

func (g *Game) AddPlayer(player *NetworkPlayer) {
	g.Lock()
	defer g.Unlock()
	if _, found := g.players[player]; found ||  g.playersInGame == 2 {
		return nil
	}
	g.players[player] = true
	player.Game = g
	g.playersInGame++
}

func (g *Game) RemovePlayer(player *NetworkPlayer) {
	g.Lock()
	defer g.Unlock()
	if _, found := g.players[player]; !found ||  g.playersInGame == 0 {
		return nil
	}
	player.disconnect()
	delete(g.players, player)
	g.playersInGame--
}

func (g *Game) PlayTurn(player *NetworkPlayer, point BoardPoint) *NetworkPlayer {
	g.board.MakeTurn(player, point)
	players := maps.Keys(g.players)
	if player == players[0] {
		g.turn = players[1]
	} else {
		g.turn = players[0]
	}
	return g.turn
}

func (g *Game) CheckWinner() *NetworkPlayer {
	return g.board.CheckWinner()
}
