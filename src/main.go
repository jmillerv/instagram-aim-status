package main

import (
	"github.com/jmillerv/instagram-aim-status/src/requests"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"strings"
)

const (
	rapid_api_app         = "RAPID_API_APP"
	rapid_api_request_url = "RAPID_API_REQUEST_URL"
	rapid_api_key         = "RAPID_API_KEY"
	rapid_api_host        = "RAPID_API_HOST"
)

func main() {
	host := os.Getenv(rapid_api_host)
	key := os.Getenv(rapid_api_key)
	url := os.Getenv(rapid_api_request_url)
	req := requests.NewClient(url, host, key)

	app := &cli.App{
		Name:    "Instagram AIM Status",
		Usage:   "CLI tool for getting 60 characters of a song",
		Version: "0.0.1",
		Author:  "Jeremiah Miller",
		Commands: cli.Commands{
			{
				Name:      "get-artist",
				Aliases:   []string{"get"},
				Usage:     "get-artist {string}",
				UsageText: "retrieves random song from the inputted artist",
				Action: func(c *cli.Context) {
					var query []string
					// get inputted string from command line
					for i := 0; i <= len(c.Args()); i++ {
						query = append(query, c.Args().Get(i))
					}
					artist := strings.Join(query, " ")
					if artist == "" {
						log.Info("no artist requested, exiting")
						os.Exit(1)
					}
					log.Infof("getting artist %s", artist)
					// query mode = input an artist string
					response, err := req.MultiSearch(artist)
					if err != nil {
						log.WithError(err).Error("unable to return artist")
					}
					log.Println(response)
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
