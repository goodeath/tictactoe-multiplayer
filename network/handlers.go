/*
 * Handlers all events inside websocket
 */
package network

import (
	"encoding/json"
	"log"
	"strconv"
)

func HandlerPlayerMove(event Event, client *NetworkPlayer) error {
	if !client.IsMyTurn() {
		client.data <- []byte("Not allowed")
		return nil
	}
	var payload PlayerMoveEvent
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		log.Println("Error decoding JSON: %v", err)
		return nil
	}
	client.PlayTurn(payload.PosX, payload.PosY)
	

	game := client.GetGame()
	players := game.GetPlayers()
	for _, player := range players {
		if player != client {
			player.data <- []byte( "{\"type\": \"player_move\", \"payload\":{\"PosX\":" + strconv.Itoa(payload.PosX) + ",\"PosY\":"+strconv.Itoa(payload.PosY)+"}}")
		}
	}

	TerminateMatch(client)
	return nil
}


func TerminateMatch(player *NetworkPlayer)  {
	game := player.GetGame()
	if winner := game.CheckWinner(); winner != nil {
		log.Println("Game Ended")
		players := game.GetPlayers()
		for _, player := range players {
			if winner == player {
				player.data <- []byte("Winner") 
			} else {
				player.data <- []byte("Loser")
			}
		}
	}
}
