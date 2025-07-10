package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const API_URL = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
const API_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MTEsImVtYWlsIjoic3NhbmNoZXpnbzAyQGdtYWlsLmNvbSIsImV4cCI6MTc1MjIxMTk2MCwiaWQiOiIiLCJwYXNzd29yZCI6ImBhbGlhc2AvKiovRlJPTS8qKi91c2Vycy8qKi9VTklPTi8qKi9TRUxFQ1QvKiovJ2FkbWluJywnYWRtaW4xMjMnLyoqLy0tIn0.bdgDhvH1KF3vhzOqW05rerSlRSrlEcF5H-0o6ELGCn8"

type StockItem struct {
	Ticker     string `json:"ticker"`
	Company    string `json:"company"`
	Brokerage  string `json:"brokerage"`
	Action     string `json:"action"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type APIResponse struct {
	Items    []StockItem `json:"items"`
	NextPage string      `json:"next_page"`
}

func fetchStockData(nextPage string) (*APIResponse, error) {
	url := API_URL
	if nextPage != "" {
		url += "?next_page=" + nextPage
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+API_TOKEN)
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data APIResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	var nextPage string
	for {
		result, err := fetchStockData(nextPage)
		if err != nil {
			log.Fatalf("Error al obtener datos: %v", err)
		}

		for _, item := range result.Items {
			fmt.Printf("Ticker: %s, Compañía: %s, Acción: %s por %s\n", item.Ticker, item.Company, item.Action, item.Brokerage)
		}

		if result.NextPage == "" {
			break // no hay más páginas
		}

		nextPage = result.NextPage
	}
}
