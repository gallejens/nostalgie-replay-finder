package main

func main() {
	// first we populate previous track map
	populateTrackList()

	// then we start polling, to check new tracks
	go startTrackPolling()

	// we start the web server
	startServer()
}
