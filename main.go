package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type APIResponse struct {
	Data []Track `json:"data"`
}

type Track struct {
	ID       int    `json:"id"`
	RadioID  int    `json:"radio_id"`
	TrackID  string `json:"track_id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	PlayedAt string `json:"played_at"`
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

	rateLimitRemaining, err := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining"))
	if err != nil || rateLimitRemaining < 200 {
		fmt.Println("Rate limit has gone below 200")
		return Track{}, err
	}

	trackResponse := APIResponse{}
	err = json.Unmarshal(body, &trackResponse)
	if err != nil {
		return Track{}, err
	}

	return trackResponse.Data[0], nil
}

func isSameTrack(first Track, second Track) bool {
	return first.Artist == second.Artist && first.Title == second.Title
}

func main() {
	var previousTrack Track = Track{}

	for {
		currentTrack, err := getCurrentTrack()
		if err != nil {
			fmt.Println(err)
			time.Sleep(30 * time.Second)
			continue
		}

		var sleepDuration time.Duration = 10

		if !isSameTrack(previousTrack, currentTrack) {
			if previousTrack != (Track{}) {
				sleepDuration = 90
			}

			fmt.Printf("Now playing: %s - %s\n", currentTrack.Title, currentTrack.Artist)
		}

		previousTrack = currentTrack
		time.Sleep(sleepDuration * time.Second)
	}
}
