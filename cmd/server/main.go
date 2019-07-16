package main

import (
	"net/http"
    "../../pkg/realtime_api"
)

func main() {
	http.HandleFunc("/api/v0/realtime", realtime_api.HandleServe)
	http.Handle("/", http.FileServer(http.Dir("./web/")))
	http.ListenAndServe(":8000", nil)
}
