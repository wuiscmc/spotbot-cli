package spotbot

import (
	"log"

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

type Playlist struct {
	tracks []Track
}
type Spotbot struct {
	rootUrl string
	fb      *firego.Firebase
}

func New(firebaseUrl string) *Spotbot {
	return &Spotbot{fb: firego.New(firebaseUrl)}
}

func (sp *Spotbot) CurrentTrack() Track {
	var val map[string]interface{}
	ref := sp.fb.Child("current_track")
	logError(ref.Value(&val))
	return toTrack(val)
}

func toTrack(val map[string]interface{}) Track {
	return Track{val["duration"].(float64), val["uri"].(string), val["title"].(string), val["image"].(string), val["artists"].(interface{})}
}

func (sp *Spotbot) Playing() Playlist {
	var val []map[string]interface{}
	ref := sp.fb.Child("playlist")
	logError(ref.Value(&val))
	tracks := make([]Track, len(val))
	for _, trackData := range val {
		tracks = append(tracks, toTrack(trackData))
	}
	playlist := Playlist{tracks}
	return playlist
}

func (sp *Spotbot) NextSong() {
	ref := sp.fb.Child("player/next")
	logError(ref.Set(true))
}

func (sp *Spotbot) Shuffle() {
	ref := sp.fb.Child("playlist/shuffle")
	shuffle := !sp.IsShuffled()
	logError(ref.Set(shuffle))
}

func (sp *Spotbot) IsShuffled() bool {
	var val bool
	ref := sp.fb.Child("playlist/shuffle")
	logError(ref.Value(&val))
	return val
}

func (sp *Spotbot) Pause() {
	ref := sp.fb.Child("player/playing")
	logError(ref.Set(false))
}

func (sp *Spotbot) Play() {
	ref := sp.fb.Child("player")
	logError(ref.Set(true))
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
