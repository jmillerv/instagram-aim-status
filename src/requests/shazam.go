package requests

import (
	"encoding/json"
	"github.com/jmillerv/instagram-aim-status/src/parser"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	searchType              = "SONGS_ARTISTS"
	header_x_rapid_api_key  = "X-RapidAPI-Key"
	header_x_rapid_api_host = "X-RapidAPI-Host"
)

// MultiSearch partially implements Shazam Core API GET MultiSearch
// this function only queries for SONGS_ARTISTS and can not do SONGS and ARTISTS independently
func (c *Client) MultiSearch(artist string) (string, error) {
	// build query url
	queryUrl := c.Url + "?search_type=" + searchType + "&query=" + artist
	log.Infof("query url: %s", queryUrl)
	// instantiate a new request
	req, err := http.NewRequest(http.MethodGet, queryUrl, nil)
	if err != nil {
		return "", err
	}

	// add headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add(header_x_rapid_api_key, c.Key)
	req.Header.Add(header_x_rapid_api_host, c.Host)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	// load body into struct
	var shazamResponse ShazamResponse

	err = json.Unmarshal(body, &shazamResponse)
	if err != nil {
		return "", err
	}
	// get characters
	characters := parser.GetFirst60Characters(shazamResponse.GetRandomTrackUrl())

	return characters, nil
}
