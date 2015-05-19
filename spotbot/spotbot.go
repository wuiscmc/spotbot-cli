package spotbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CloudCom/firego"
)

type Track struct {
	duration float64
	uri      string
	title    string
	artist   string
}

func (track Track) String() string {
	res := fmt.Sprintf("%s  -  %s", track.title, track.artist)
	return res
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
	artist := val["artists"].([]interface{})[0].(string)
	return Track{val["duration"].(float64), val["uri"].(string), val["title"].(string), artist}
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

func (sp *Spotbot) NextSong() {
	requestServer("next")
}

func (sp *Spotbot) Pause() {
	requestServer("pause")
}

func (sp *Spotbot) Play() {
	requestServer("play")
}

func requestServer(action string) {
	url := "http://office-robot.local:3030/"
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url+action, nil)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println("   ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)

		}
		fmt.Println(contents)

	}

}
func (sp *Spotbot) Search(query string) []Track {
	if query == "" {
		return nil
	}
	url := fmt.Sprintf("http://api.spotify.com/v1/search?limit=20&type=track&market=se&q='%s'", query)
	fmt.Println(url)
	res, _ := http.Get(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var data map[string]map[string]interface{}
	json.Unmarshal(body, &data)
	if data != nil {
		rawItems := data["tracks"]["items"].([]interface{})
		tracks := make([]Track, 0)
		for _, rawItem := range rawItems {
			item := rawItem.(map[string]interface{})
			artist := item["artists"].([]interface{})[0].(map[string]interface{})["name"].(string)
			track := Track{title: item["name"].(string), uri: item["uri"].(string), artist: artist}
			tracks = append(tracks, track)
		}
		return tracks
	} else {
		return nil
	}
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
