package realtime_api

import (
    "./client"
    "./server"
    "testing"
    "net/http/httptest"
    "net/http"
)

func startClientServer() (*client.Client, error) {
    s := httptest.NewServer(http.HandlerFunc(server.HandleServe))
    defer s.Close()
    url := "ws" + s.URL[4:]
    c, err := client.NewClient(url)
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