package server

import (
    "../common"
    "../../status_codes"
    "../../keys/keychain"
    "encoding/json"
    "github.com/gorilla/websocket"
    "./methods"
    "net/http"
    "log"
    "time"
)

func HandleServe(w http.ResponseWriter, r *http.Request) {
    // Define check to allow websocket connection
    upgrader := websocket.Upgrader{
        CheckOrigin: func(r *http.Request) (bool) {
            return true
        },
    }
    // Convert request to websocket and quick if error
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    // Always close websocket when done
    defer ws.Close()

    for {
        // Read Message from Websocket
        _, raw, err := ws.ReadMessage()
        if err != nil {
            log.Println(err)
            break
        }

        // Create Basic Reply Message
        var reply common.OutgoingMessage
        reply.Type = "result"
        reply.Result = status_codes.INTERNAL_SERVER_ERROR

        // Check if is Valid JSON and convert to incommingMessage structure
        var msg common.IncommingMessage
        err = json.Unmarshal(raw, &msg)
        if err != nil {
            log.Println(err)
        }

        // Assume the same ID
        reply.Id = msg.Id

        // If has a valid message ID
        if msg.Type != "" && msg.Id != 0 {
            // Handle Different Types of Messages
            switch msg.Type {
            case "method":
                // If method name exists, use it
                if m, ok := methods.Get(msg.Method); ok {
                    reply.Obj, reply.Result = m(msg.Obj)
                } else {
                    reply.Result = status_codes.BAD_REQUEST
                }
            case "lookup":
                log.Println("Lookup")
                reply.Result = status_codes.OK
            case "subscribe":
                for _, id := range msg.Obj {
                    switch id.(type) {
                    case float64:
                        // Retrieve key, return 404 if no exist, and append key to output array
                        k, err := keychain.GetKey(uint64(id.(float64)))
                        reply.Result = status_codes.ACCEPTED
                        if err == nil {
                            go func(id uint64) {
                                var subMsg common.OutgoingMessage
                                subMsg.Type = "update"
                                subMsg.Id = msg.Id
                                for {
                                    token, _ := k.GetCode()
                                    subMsg.Obj = []interface{}{
                                        struct{Token string `json:"token"`}{
                                            Token: token,
                                        },
                                    }
                                    subMsgRaw, _ := json.Marshal(&subMsg)
                                    ws.WriteMessage(websocket.TextMessage, subMsgRaw)
                                    time.Sleep(30*time.Second)
                                }
                            }(msg.Id)
                        } else {
                            reply.Result = status_codes.NOT_FOUND
                        }
                    default:
                        reply.Result = status_codes.BAD_REQUEST
                    }
                }
            }
        } else {
            reply.Result = status_codes.BAD_REQUEST
        }

        // Send Reply
        replyRaw, _ := json.Marshal(&reply)
        ws.WriteMessage(websocket.TextMessage, replyRaw)
    }
}
