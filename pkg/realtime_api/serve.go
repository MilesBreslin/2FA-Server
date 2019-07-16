package realtime_api

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    "net/http"
    "log"
)

func HandleServe(w http.ResponseWriter, r *http.Request) {
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
        _, raw, err := ws.ReadMessage()
        if err != nil {
            log.Println(err)
            break
        }

        var msg incommingMessage
        err = json.Unmarshal(raw, &msg)
        if err != nil {
            log.Println(err)
        }

        log.Println(msg)
        switch msg.Type {
        case "subscribe":
            log.Println("Subscribe")
            var reply outgoingMessage
            reply.Id = msg.Id
            reply.Result = 202
            replyRaw, _ := json.Marshal(&reply)
            ws.WriteMessage(1, replyRaw)
        }
    }
}