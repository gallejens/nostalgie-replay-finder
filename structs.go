package main

type Track struct {
	Title         string `json:"title"`
	Artist        string `json:"artist"`
	PlayedAt      string `json:"playedAt"`
	AlreadyPlayed bool   `json:"alreadyPlayed"`
}

type WebSocketKeepAlive struct {
	KeepAlive bool `json:"keepAlive"`
}
