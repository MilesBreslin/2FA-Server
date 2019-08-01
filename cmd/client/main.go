package main

import (
    "../../pkg/realtime_api/client"
    "time"
    "log"
)

func main() {
    c, err := client.NewClient("ws://localhost:8000/api/v0/realtime")
    if err != nil {
    	log.Println(err)
    }
    arr := make([]interface{}, 1)
    arr[0] = "hi"
    r := c.Test(arr)
    log.Println(r)
    time.Sleep(10 * time.Second)
}
