package client

import (
	"../common"
	"github.com/gorilla/websocket"
	"log"
)

type client struct {
	ws				websocket.Conn
}

func NewClient(url string) (*client, error) {
	var c client
	// url example : "ws://localhost:8000/api/v0/realtime"
	c.ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Error initializing connection")
	}
	go handleClient()
	return &c, nil
}

func handleClient() {

}