package main

import (
	"fmt"
	"os"

	"github.com/cosn/firebase"
)

type CurrentTrack struct {
	duration string
	uri      string
	title    string
	image    string
	artists  interface{}
}

type Spotbot struct {
	firebase *firebase.Client
}

func (sp *Spotbot) CurrentTrack() string {
	value := sp.firebase.Child("current_track", nil, nil).Value()
	val := value.(map[string]interface{})
	return val["title"].(string)
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
