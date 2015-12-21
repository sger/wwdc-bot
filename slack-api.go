package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
)

const URL_SLACK_API = "https://api.slack.com/"
const URL_SLACK_API_RTM_START = "https://slack.com/api/rtm.start"

var counter uint64

type RtmStartResponse struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	Url   string       `json:"url"`
	Self  SelfResponse `json:"self"`
}

type SelfResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	User    string `json:"user"`
}

func GetMessage(ws *websocket.Conn) (m Message, err error) {
	err = websocket.JSON.Receive(ws, &m)
	return
}

func PostMessage(ws *websocket.Conn, m Message) error {
	m.Id = atomic.AddUint64(&counter, 1)
	return websocket.JSON.Send(ws, m)
}

func Connect(token string) (*websocket.Conn, *RtmStartResponse, error) {
	r, err := Start(token)
	if err != nil {
		log.Fatal(err)
	}

	ws, err := websocket.Dial(r.Url, "", URL_SLACK_API)
	if err != nil {
		log.Fatal(err)
	}

	return ws, r, err
}

func Start(token string) (*RtmStartResponse, error) {

	url := fmt.Sprintf(URL_SLACK_API_RTM_START+"?token=%s", token)
	resp, err := http.Get(url)

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

	var result RtmStartResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
