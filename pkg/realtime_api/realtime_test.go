package realtime_api

import (
    "./client"
    "./server"
    "testing"
    "net/http"
    "time"
)

func startClientServer() (*client.Client, error) {
    http.HandleFunc("/realtime", server.HandleServe)
    go http.ListenAndServe(":8000", nil)
    time.Sleep(5 * time.Second)
    c, err := client.NewClient("ws://127.0.0.1:8000/realtime")
    return c, err
}

func TestTestFunc(t *testing.T) {
    c, err := startClientServer()
    if err != nil {
        t.Errorf("%v",err)
        return
    }
    arr := make([]interface{}, 1)
    arr[0] = "hi"
    r := c.Test(arr)
    if r[0] != arr[0] {
        t.Errorf("Expecting test to receive %v, but got %v", arr, r)
    }
}