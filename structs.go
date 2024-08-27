package main

type Track struct {
	ID       int    `json:"id"`
	RadioID  int    `json:"radio_id"`
	TrackID  string `json:"track_id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	PlayedAt string `json:"played_at"`
}
