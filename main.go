package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wwdc-bot token\n @bot-name uiview or @bot-name 101 2015")
		os.Exit(1)
	}

	ws, r, err := Connect(os.Args[1])
	fmt.Println("wwdc-bot is running hit ^C to exit")

	if err != nil {
		log.Fatal(err)
	}

	for {
		m, err := GetMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+r.Self.Id+">") {
			fields := strings.Fields(m.Text)
			if len(fields) == 3 {
				go func(m Message) {

					resp, err := getSession(fields[1], fields[2])
					if err != nil {
						m.Text = "no results"
						PostMessage(ws, m)
					} else {
						downloadUrl := fmt.Sprintf("https://developer.apple.com/videos/wwdc/%.0f/?id=%.0f", resp.Year, resp.Number)
						m.Text = "Title: " + resp.Title + "\nDescription: " + resp.Description + "\nDownload url: " + downloadUrl
						PostMessage(ws, m)
					}

				}(m)
			} else if len(fields) == 2 {
				go func(m Message) {

					resp, err := search(fields[1])
					if err != nil {
						log.Fatal(err)
					}
					var result string

					if len(resp.Results) == 0 {
						m.Text = "no results"
						PostMessage(ws, m)
					} else {
						for _, element := range resp.Results[0:4] {
							downloadUrl := fmt.Sprintf("https://developer.apple.com/videos/wwdc/%.0f/?id=%.0f", element.Year, element.Number)
							result += element.Title + "\n" + element.Description + "\n" + downloadUrl + "\n\n"
						}
						m.Text = result
						PostMessage(ws, m)
					}
				}(m)
			} else {
				go func(m Message) {
					m.Text = "sorry " + "<@" + m.User + ">" + " wrong usage please write @" + r.Self.Name + ": 101 2015 or search @" + r.Self.Name + ": uiview \n 101 is the id of session and 2015 the year\nmore info here http://asciiwwdc.com"
					PostMessage(ws, m)
				}(m)
			}
		}
	}
}
