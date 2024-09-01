package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type APIResponse struct {
	Data []APIResponseDataEntry `json:"data"`
}

type APIResponseDataEntry struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	PlayedAt string `json:"played_at"`
}

func getFromNostalgieAPI() (APIResponse, error) {
	resp, err := http.Get(API_ENDPOINT)
	if err != nil {
		return APIResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, err
	}

	rateLimitRemaining, err := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining"))
	// fmt.Println("Rate limit remaining:", rateLimitRemaining)
	if err != nil || rateLimitRemaining < 200 {
		fmt.Println("Rate limit has gone below 100, watch out")
		return APIResponse{}, err
	}

	response := APIResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return APIResponse{}, err
	}

	return response, nil
}
