package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TracksResponse struct {
	Data []Track `json:"data"`
}

type Track = struct {
	ID                int    `json:"id"`
	RadioID           int    `json:"radio_id"`
	TrackID           string `json:"track_id"`
	Title             string `json:"title"`
	Artist            string `json:"artist"`
	DurationInSeconds int    `json:"duration_in_seconds"`
	PlayedAt          string `json:"played_at"`
}

func getCurrentTrack() (Track, error) {
	resp, err := http.Get("https://mediahuis.pointbreak.dev/api/v1/radio/9/tracks")
	if err != nil {
		return Track{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Track{}, err
	}

	trackResponse := TracksResponse{}
	err = json.Unmarshal(body, &trackResponse)
	if err != nil {
		return Track{}, err
	}

	return trackResponse.Data[0], nil
}

func main() {
	currentTrack, err := getCurrentTrack()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentTrack.Title)
}
