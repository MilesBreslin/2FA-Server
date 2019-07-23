package main

import (
	"net/http"
    "../../pkg/realtime_api"
    _ "../../pkg/static_web"
)

func main() {
	http.HandleFunc("/api/v0/realtime", realtime_api.HandleServe)
	http.Handle("/", http.FileServer(http.Dir("./web/")))
	http.ListenAndServe(":8000", nil)
}
