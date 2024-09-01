package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Track)

func startServer() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/initial", handleInitial)
	http.HandleFunc("/ws", handleConnections)

	go broadcastMessages()

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/dist/index.html")
}

func handleInitial(w http.ResponseWriter, r *http.Request) {
	jsonTracklist, err := json.Marshal(getInitialTracklist())
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTracklist)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg WebSocketKeepAlive
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(clients, conn)
			return
		}
	}
}

func broadcastMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func addMessageToBroadcast(track Track) {
	broadcast <- track
}
