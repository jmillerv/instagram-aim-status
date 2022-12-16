package requests

import (
	"math/rand"
	"time"
)

// ShazamResponse is a Json struct of the data returned from the MultiSearch SONGS_ARTIST search type
type ShazamResponse struct {
	Tracks struct {
		Hits []struct {
			Track struct {
				Layout   string `json:"layout"`
				Type     string `json:"type"`
				Key      string `json:"key"`
				Title    string `json:"title"`
				Subtitle string `json:"subtitle"`
				Share    struct {
					Subject  string `json:"subject"`
					Text     string `json:"text"`
					Href     string `json:"href"`
					Image    string `json:"image"`
					Twitter  string `json:"twitter"`
					HTML     string `json:"html"`
					Avatar   string `json:"avatar"`
					Snapchat string `json:"snapchat"`
				} `json:"share"`
				Images struct {
					Background string `json:"background"`
					Coverart   string `json:"coverart"`
					Coverarthq string `json:"coverarthq"`
					Joecolor   string `json:"joecolor"`
				} `json:"images"`
				Hub struct {
					Type    string `json:"type"`
					Image   string `json:"image"`
					Actions []struct {
						Name string `json:"name"`
						Type string `json:"type"`
						ID   string `json:"id,omitempty"`
						URI  string `json:"uri,omitempty"`
					} `json:"actions"`
					Options []struct {
						Caption string `json:"caption"`
						Actions []struct {
							Name string `json:"name"`
							Type string `json:"type"`
							URI  string `json:"uri"`
							ID   string `json:"id,omitempty"`
						} `json:"actions"`
						Beacondata struct {
							Type         string `json:"type"`
							Providername string `json:"providername"`
						} `json:"beacondata"`
						Image               string `json:"image"`
						Type                string `json:"type"`
						Listcaption         string `json:"listcaption"`
						Overflowimage       string `json:"overflowimage"`
						Colouroverflowimage bool   `json:"colouroverflowimage"`
						Providername        string `json:"providername"`
					} `json:"options"`
					Providers []struct {
						Caption string `json:"caption"`
						Images  struct {
							Overflow string `json:"overflow"`
							Default  string `json:"default"`
						} `json:"images"`
						Actions []struct {
							Name string `json:"name"`
							Type string `json:"type"`
							URI  string `json:"uri"`
						} `json:"actions"`
						Type string `json:"type"`
					} `json:"providers"`
					Explicit    bool   `json:"explicit"`
					Displayname string `json:"displayname"`
				} `json:"hub"`
				Artists []struct {
					ID     string `json:"id"`
					Adamid string `json:"adamid"`
				} `json:"artists"`
				URL string `json:"url"`
			} `json:"track"`
		} `json:"hits"`
	} `json:"tracks"`
	Artists struct {
		Hits []struct {
			Artist struct {
				Avatar   string `json:"avatar"`
				Name     string `json:"name"`
				Verified bool   `json:"verified"`
				Weburl   string `json:"weburl"`
				Adamid   string `json:"adamid"`
			} `json:"artist"`
		} `json:"hits"`
	} `json:"artists"`
}

// GetRandomTrackUrl selects one of the tracks from the response and returns its url
func (s *ShazamResponse) GetRandomTrackUrl() string {
	rand.Seed(time.Now().UnixNano())
	key := rand.Intn(len(s.Tracks.Hits)) + 1 // I copied this from the last time I needed to create a random key and can't remember what the +1 fixes and didn't comment it in my old code. Cargo culting myself and you here.
	return s.Tracks.Hits[key].Track.URL
}
