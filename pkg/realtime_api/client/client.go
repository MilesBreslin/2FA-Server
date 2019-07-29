package client

import (
    "github.com/opensource2fa/pkg/realtime_api/common"
    "github.com/gorilla/websocket"
    "log"
    "errors"
    "encoding/json"
)

type Client struct {
    ws              *websocket.Conn
    send            chan interface{}
    quit            chan struct{}
    methodAdd       chan methodAdd
    methodReceive   chan *common.OutgoingMessage
    methodsRunning  map[uint64] chan common.OutgoingMessage
    newUID          chan uint64
}

type methodAdd struct {
    result          chan common.OutgoingMessage
    id              uint64
}

func NewClient(url string) (*Client, error) {
    var c Client
    // url example : "ws://localhost:8000/api/v0/realtime"
    ws, _, err := websocket.DefaultDialer.Dial(url, nil)
    c.ws = ws
    if err != nil {
        log.Println(err)
        return nil, errors.New("Error initializing connection")
    }

    c.send = make(chan interface{}, 0)
    c.quit = make(chan struct{}, 0)
    c.methodAdd = make(chan methodAdd, 0)
    c.methodReceive = make(chan *common.OutgoingMessage, 0)
    c.newUID = make(chan uint64, 5)
    c.methodsRunning = make(map[uint64] chan common.OutgoingMessage)

    go c.handleClient()
    return &c, nil
}

func (c *Client) handleClient() {
    defer c.ws.Close()

    // Generate New UIDs
    go func() {
        for i := uint64(1); ; i++ {
            c.newUID <- i
        }
    }()

    // Method Handler Thread
    go func() {
        for {
            select {
            case m := <- c.methodAdd:
                c.methodsRunning[m.id] = m.result
            case r := <- c.methodReceive:
                methodId := r.Id
                c.methodsRunning[methodId] <- *r
                close(c.methodsRunning[methodId])
                delete(c.methodsRunning,methodId)
            }
        }
    }()

    // Read Thread
    go func() {
        for {
            _, raw, err := c.ws.ReadMessage()
            if err != nil {
                if websocket.IsCloseError(err) {
                    return
                }
                log.Println(err)
                continue
            }
            var msg common.OutgoingMessage
            err = json.Unmarshal(raw, &msg)
            if err != nil {
                log.Println(err)
                continue
            }
            switch msg.Type {
            case "result":
                c.methodReceive <- &msg
            default:
                log.Println("Unknown server message\n%s", string(raw))
            }
        }
    }()

    // Write Thread
    go func() {
        for {
            obj := <-c.send
            msg, err := json.Marshal(obj)
            err = c.ws.WriteMessage(websocket.TextMessage, msg)
            if err != nil {
                log.Println(err)
                return
            }
        }
    }()

    <-c.quit
}

func (c *Client) getUID() (uint64) {
    return <- c.newUID
}

func (c *Client) runMethod(method string, obj []interface{}) ([]interface{}, uint16) {
    var send common.IncommingMessage
    send.Type = "method"
    send.Method = method
    send.Obj = obj
    send.Id = c.getUID()

    methodRequest := methodAdd{
        result: make(chan common.OutgoingMessage,0),
        id: send.Id,
    }
    c.methodAdd <- methodRequest
    c.send <- send

    msg := <- methodRequest.result
    return msg.Obj, msg.Result
}