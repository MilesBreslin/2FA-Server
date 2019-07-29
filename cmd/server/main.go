package main

import (
	"net/http"
    "github.com/opensource2fa/pkg/realtime_api/server"
    _ "github.com/opensource2fa/pkg/static_web"
)

func main() {
	http.HandleFunc("/api/v0/realtime", server.HandleServe)
	http.ListenAndServe(":8000", nil)
}
