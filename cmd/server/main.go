package main

import (
	"net/http"
	"log"
    "github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/api/v0/realtime", handleWebsocket)
	http.Handle("/", http.FileServer(http.Dir("./web/")))
	http.ListenAndServe(":8000", nil)
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) (bool) {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	defer ws.Close()

	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		ws.WriteMessage(messageType, msg)
	}
}