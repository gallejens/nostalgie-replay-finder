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
	response, err := getFromNostalgieAPI()
	if err != nil {
		log.Fatal(err)
	}

	// nostalgie is returning the tracks in reverse order for our purposes
	slices.Reverse(response.Data)

	for _, track := range response.Data {
		tracklist = append(tracklist, Track{
			Title:         track.Title,
			Artist:        track.Artist,
			PlayedAt:      track.PlayedAt,
			AlreadyPlayed: checkTrackAlreadyPlayed(track),
		})
	}

	fmt.Println("Tracklist has been populated with previous tracks")
}

func startTrackPolling() {
	for {
		response, err := getFromNostalgieAPI()
		if err != nil {
			fmt.Println(err)
			time.Sleep(30 * time.Second)
			continue
		}

		lastDataEntry := response.Data[0]

		var previousTrack = tracklist[len(tracklist)-1]
		var sleepDuration time.Duration = POLLING_RATE

		if previousTrack.Artist != lastDataEntry.Artist || previousTrack.Title != lastDataEntry.Title {
			sleepDuration = POLLING_TIMEOUT_ON_MATCH

			newTrack := Track{
				Title:         lastDataEntry.Title,
				Artist:        lastDataEntry.Artist,
				PlayedAt:      lastDataEntry.PlayedAt,
				AlreadyPlayed: checkTrackAlreadyPlayed(lastDataEntry),
			}

			fmt.Printf("Now playing: %s - %s | Already played: %s\n", newTrack.Title, newTrack.Artist, strconv.FormatBool(newTrack.AlreadyPlayed))
			addMessageToBroadcast(newTrack)
			tracklist = append(tracklist, newTrack)
		}

		time.Sleep(sleepDuration * time.Second)
	}
}

func checkTrackAlreadyPlayed(track APIResponseDataEntry) bool {
	for _, t := range tracklist {
		if t.Title == track.Title && t.Artist == track.Artist {
			return true
		}
	}

	return false
}

func getInitialTracklist() []Track {
	return tracklist
}
