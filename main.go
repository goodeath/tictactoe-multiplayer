package main

import (
	"tictactoe/websocket"
	"tictactoe/network"
)

func main() {
	server := network.NewServer()
	ws := websocket.NewWSManager(server.HandleConnection);
	ws.StartServer()
}
