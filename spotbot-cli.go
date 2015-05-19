package main

import (
	"fmt"
	"os"

	"github.com/wuiscmc/spotbot-cli/spotbot"
)

func main() {
	firebaseUrl := os.Getenv("FIREBASE_URL")
	if firebaseUrl == "" {
		fmt.Println("Please set up your FIREBASE_URL env variable first")
		return
	}
	sp := spotbot.New(firebaseUrl)
	fmt.Println("%s", sp.CurrentTrack())
	sp.Playing()
}
