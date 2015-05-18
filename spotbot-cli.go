package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CloudCom/firego"
)

type Track struct {
	duration float64
	uri      string
	title    string
	image    string
	artists  interface{}
}

type Player struct {
	next    bool
	playing bool
}

type Spotbot struct {
	rootUrl string
	fb      *firego.Firebase
}

func (sp *Spotbot) CurrentTrack() Track {
	sp.setRef("current_track")
	var val map[string]interface{}
	if err := sp.fb.Value(&val); err != nil {
		log.Fatal(err)

	}
	track := Track{val["duration"].(float64), val["uri"].(string), val["title"].(string), val["image"].(string), val["artists"].(interface{})}
	return track
}

func (sp *Spotbot) NextSong() {
	sp.setRef("player")
	v := map[string]bool{"next": true}
	if err := sp.fb.Set(v); err != nil {
		log.Fatal(err)
	}
}

func (sp *Spotbot) setRef(url string) {
	fb := firego.New(os.Getenv("FIREBASE_URL") + "/" + url)
	sp.fb = fb
}

func main() {
	sp := &Spotbot{}
	fmt.Println("%s", sp.CurrentTrack())
	sp.NextSong()
}
