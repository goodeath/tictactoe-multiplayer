package network

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
)

var routes map[string]EventHandler = map[string]EventHandler {
	PlayerMoveMessage: HandlerPlayerMove,
}

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

func Router(player *NetworkPlayer){
	go listen(player)
	go talk(player)
}


func listen(player *NetworkPlayer){
	log.Println("Listening")
	defer player.disconnect()
	for {
		_, payload, err := player.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		var request Event

		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling message: %v", err)
			break
		}

		if err := RouteEvent(request, player); err != nil {
			log.Println("Error handeling Message: ", err)
		}
	}
}


func RouteEvent(event Event, player *NetworkPlayer) error {
	if handler, ok := routes[event.Type]; ok {
		if err := handler(event, player); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}


func talk(player *NetworkPlayer) {
	defer player.disconnect()
	for {
		select {
		case message, ok := <-player.data:
			if !ok {
				if err := player.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return 
			}
			if err := player.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println(err)
			}
			log.Println("sent message")
		}
	}
}
