package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"time"
)

var tracklist []Track

func populateTrackList() {
	tracks, err := getTracksFromNostalgieApi()
	if err != nil {
		log.Fatal(err)
	}

	slices.Reverse(tracks)
	tracklist = tracks

	fmt.Println("Tracklist has been populated with previous tracks. Starting application")
}

func startTrackPolling() {
	for {
		tracks, err := getTracksFromNostalgieApi()
		if err != nil {
			fmt.Println(err)
			time.Sleep(30 * time.Second)
			continue
		}

		newTrack := tracks[0]

		var sleepDuration time.Duration = POLLING_RATE

		if !isSameTrack(tracklist[len(tracklist)-1], newTrack) {
			sleepDuration = 90

			alreadyPlayed := checkTrackAlreadyPlayed(newTrack)
			logMessage := fmt.Sprintf("Now playing: %s - %s | Already played: %s", newTrack.Title, newTrack.Artist, strconv.FormatBool(alreadyPlayed))

			fmt.Println(logMessage)
			broadcastToWeb(logMessage)
		}

		tracklist = append(tracklist, newTrack)
		time.Sleep(sleepDuration * time.Second)
	}
}

func checkTrackAlreadyPlayed(track Track) bool {
	for _, t := range tracklist {
		if isSameTrack(t, track) {
			return true
		}
	}

	return false
}

func main() {
	// first we populate previous track map
	populateTrackList()

	// then we start polling, to check new tracks
	go func() {
		startTrackPolling()
	}()

	// we start the web server
	startServer()
}
