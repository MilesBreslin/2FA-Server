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

func TestKeys(t *testing.T) {
    c, err := startClientServer()
    if err != nil {
        t.Errorf("%v",err)
        return
    }
    key, err := c.GetKey(0)
    if err == nil {
        t.Errorf("Tried to get key 0 before every adding any keys and got no error")
    }
    if key != nil {
        t.Errorf("Returned key pointer is not nil, when expecting empty pointer")
    }
    exampleSecret := "lhe4kfhfqapxipzmohswb6i5adg2gauh"
    id, err := c.AddKey(exampleSecret)
    if err != nil {
        t.Errorf("Error adding initial key: %v", err)
    }
    if id == 0 {
        t.Errorf("First added key id is 0, expecting 1 or higher")
    }
    key, err = c.GetKey(id)
    if err != nil {
        t.Errorf("Error retrieving first added key: %v", err)
    }

}