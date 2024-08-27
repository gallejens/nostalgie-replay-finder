package main

func isSameTrack(first Track, second Track) bool {
	return first.Artist == second.Artist && first.Title == second.Title
}
