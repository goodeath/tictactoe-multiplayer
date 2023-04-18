package network

import "encoding/json"

type Event struct {
	Type string `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, player *NetworkPlayer) error

type PlayerMoveEvent struct {
	PosX int `json:"PosX"`
	PosY int `json:"PosY"`
}
