package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL_ASCII_WWDC_API = "http://asciiwwdc.com"

type SessionResponse struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Year        float64 `json:"year"`
	Number      float64 `json:"number"`
}

type SearchResponse struct {
	Results []ResultsData `json:"results"`
}

type ResultsData struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Year        float64 `json:"year"`
	Number      float64 `json:"number"`
}

func getSession(sessionID string, year string) (*SessionResponse, error) {

	url := fmt.Sprintf(URL_ASCII_WWDC_API+"/%s/sessions/%s", year, sessionID)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", `application/json`)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result SessionResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func search(query string) (*SearchResponse, error) {
	url := fmt.Sprintf(URL_ASCII_WWDC_API+"/search?q=%s", query)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", `application/json`)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result SearchResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
