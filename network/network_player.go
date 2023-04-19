package network

import "github.com/gorilla/websocket"

type NetworkPlayer struct {
	connection *websocket.Conn
	server *Server
	game *Game
	name string
	data chan []byte
}

func NewNetworkPlayer(server *Server, conn *websocket.Conn) *NetworkPlayer {
	player := &NetworkPlayer{
		server: server,
		connection: conn,
		data: make(chan []byte),
	}
	return player
}

func (player *NetworkPlayer) ReadMessage() (int, []byte, error) {
	return player.connection.ReadMessage()
}

func (player *NetworkPlayer) IsMyTurn() bool {
	return player.game.turn == player
}

func (player *NetworkPlayer) SetGame(game *Game) {
	player.game = game
}

func (player *NetworkPlayer) SetServer(server *Server) {
	player.server = server
}

func (player *NetworkPlayer) GetGame() *Game {
	return player.game
}

func (player *NetworkPlayer) PlayTurn(posx int, posy int) {
	player.game.PlayTurn(posx, posy)
}

func (player *NetworkPlayer) IsMatchStarted() bool {
	return !player.game.IsAvailable() 
}

func (player *NetworkPlayer) disconnect() {
	player.server.Disconnect(player)
}
