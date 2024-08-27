package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type APIResponse struct {
	Data []Track `json:"data"`
}

func getTracksFromNostalgieApi() ([]Track, error) {
	resp, err := http.Get(API_ENDPOINT)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rateLimitRemaining, err := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining"))
	// fmt.Println("Rate limit remaining:", rateLimitRemaining)
	if err != nil || rateLimitRemaining < 200 {
		fmt.Println("Rate limit has gone below 100, watch out")
		return nil, err
	}

	trackResponse := APIResponse{}
	err = json.Unmarshal(body, &trackResponse)
	if err != nil {
		return nil, err
	}

	return trackResponse.Data, nil
}
