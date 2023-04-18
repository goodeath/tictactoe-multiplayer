package network


import (
	"sync"
	"github.com/gorilla/websocket"
)

type Server struct {
	games []*Game
	//playerGame map[*NewtorkPlayer]*Game
	//connectedPlayers []*NetworkPlayer
	sync.RWMutex
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) HandleConnection(ws *websocket.Conn){
	s.Lock()
	defer s.Unlock()
	player := NewNetworkPlayer(s, ws)
	s.AddPlayer(player, nil)
	go Router(player)
}


func (s *Server) FindEmptyGame() *Game {
	for _, game := range s.games {
		if game.IsAvailable() {
			return game
		}
	}
	new_game := s.StartRoomGame()
	return new_game
}

func (s *Server) StartRoomGame() *Game {
	game := NewGame()
	s.games = append(s.games, game)
	return game
}

func (s *Server) AddPlayer(player *NetworkPlayer, game *Game) {
	if game == nil {
		game = s.FindEmptyGame()
	}
	//s.connectedPlayers.append(player)
	game.AddPlayer(player)
}


func (s *Server) RemovePlayer(player *NetworkPlayer, game *Game) {
	//game.RemovePlayer(player)
}

func (s *Server) CloseRoomGame(game *Game) { 
//
}

func (s *Server) Disconnect(player *NetworkPlayer){
}
