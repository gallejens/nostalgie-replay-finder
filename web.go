package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/r3labs/sse/v2"
)

var server *sse.Server

func startServer() {
	server = sse.New()
	server.CreateStream("messages")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/tracks", server.ServeHTTP)

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func broadcastToWeb(message string) {
	server.Publish("messages", &sse.Event{
		Data: []byte(message),
	})
}
