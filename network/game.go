// Handle game through network
package network

import (
	"sync"
	"golang.org/x/exp/maps"
)

type Game struct {
	playersInGame int
	players map[*NetworkPlayer]bool
	board *NetworkBoard
	turn *NetworkPlayer
	winner *NetworkPlayer
	sync.RWMutex
}

func NewGame() *Game{
	return &Game{
		board: NewNetworkBoard(),
		players: make(map[*NetworkPlayer]bool),
	}
}

func (g *Game) IsAvailable() bool {
	return g.playersInGame < 2
}

func (g *Game) AddPlayer(player *NetworkPlayer) {
	g.Lock()
	defer g.Unlock()
	if _, found := g.players[player]; found ||  !g.IsAvailable() {
		return
	}
	if g.turn == nil {
		g.turn = player
	}
	g.players[player] = true
	g.board.AddPlayer(player)
	player.SetGame(g)
	g.playersInGame++
}

func (g *Game) RemovePlayer(player *NetworkPlayer) {
	g.Lock()
	defer g.Unlock()
	if _, found := g.players[player]; !found ||  g.playersInGame == 0 {
		return
	}
	delete(g.players, player)
	g.board.RemovePlayer(player)
	g.playersInGame--
}

// Fix this
func (g *Game) PlayTurn(posx int, posy int) *NetworkPlayer {
	players := maps.Keys(g.players)
	player := g.turn
	g.board.PlayTurn(player, posx, posy)
	if player == players[0]{
		g.turn = players[1]
	} else {
		g.turn = players[0]
	}
	return g.turn
}

func (g *Game) GetPlayers() []*NetworkPlayer {
	players := maps.Keys(g.players)
	return players
}

func (g *Game) CheckWinner() *NetworkPlayer {
	return g.board.CheckWinner()
}
