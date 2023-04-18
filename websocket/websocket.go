package websocket

import (
	"errors"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type WSHandler func(connection *websocket.Conn)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}


type WSManager struct {
	handle WSHandler
}

func NewWSManager(handler WSHandler) *WSManager {
	return &WSManager{
		handle: handler,
	}
}


func (m *WSManager) serveWS (writer http.ResponseWriter, request *http.Request) {
	connection, err := websocketUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	m.handle(connection)
}

func (m *WSManager) StartServer(){
	log.Println("Running server")
	http.HandleFunc("/ws", m.serveWS)
	http.ListenAndServe(":8080", nil)
}

