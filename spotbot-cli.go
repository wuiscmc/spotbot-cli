package main

import (
	"fmt"
	"os"

	"github.com/cosn/firebase"
)

type Track struct {
	duration float64
	uri      string
	title    string
	image    string
	artists  interface{}
}

type Spotbot struct {
	firebase *firebase.Client
}

func (sp *Spotbot) CurrentTrack() Track {
	value := sp.firebase.Child("current_track", nil, nil).Value()
	val := value.(map[string]interface{})
	track := Track{val["duration"].(float64), val["uri"].(string), val["title"].(string), val["image"].(string), val["artists"].(interface{})}
	return track
}

func New() *Spotbot {
	fb := new(firebase.Client)
	fb.Init(os.Getenv("FIREBASE_URL"), "", nil)
	return &Spotbot{fb}
}

func main() {
	sp := New()
	fmt.Println("%s", sp.CurrentTrack())
}
