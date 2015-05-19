package main

import (
	"fmt"
	"os"

	"github.com/wuiscmc/spotbot-cli/spotbot"
)

func control(option string, sp *spotbot.Spotbot, opts interface{}) {
	switch option {
	case "play":
		sp.Play()
	case "pause":
		sp.Pause()
	case "next":
		sp.NextSong()
	case "playlist":
		//fmt.Println(sp.CurrentPlaylist())
	case "current":
		fmt.Println(sp.CurrentTrack())
	case "search":
		res := sp.Search(opts.(string))
		for _, track := range res {
			fmt.Println(track)
		}
	}
}

func main() {
	firebaseUrl := os.Getenv("FIREBASE_URL")
	if firebaseUrl == "" {
		fmt.Println("Please set up your FIREBASE_URL env variable first")
		return
	}

	sp := spotbot.New(firebaseUrl)

	var option, query string
	numArgs := len(os.Args)
	switch {
	default:
		query = ""
	case numArgs == 1:
		option = "current"
	case numArgs == 2:
		option = os.Args[1]
	case numArgs == 3:
		option = os.Args[1]
		query = os.Args[2]
	}

	control(option, sp, query)
}
