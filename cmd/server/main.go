package main

import (
	"net/http"
    "../../pkg/realtime_api/server"
    _ "../../pkg/static_web"
)

func main() {
	http.HandleFunc("/api/v0/realtime", server.HandleServe)
	http.ListenAndServe(":8000", nil)
}
