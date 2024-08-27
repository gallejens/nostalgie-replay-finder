package main

import (
	"fmt"
	"time"
)

func main() {
	// first we populate previous track map

	var previousTrack Track = Track{}

	for {
		tracks, err := apiGetTracks()
		if err != nil {
			fmt.Println(err)
			time.Sleep(30 * time.Second)
			continue
		}

		currentTrack := tracks[0]
		var sleepDuration time.Duration = POLLING_RATE

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
