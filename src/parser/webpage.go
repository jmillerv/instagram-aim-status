package parser

import (
	"context"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"strings"
)

// getLyricsFromWebPage returns the lyrics from a shazam url
func getLyricsFromWebPage(shazamUrl string) string {
	lyrics := scrapePageForLyrics(shazamUrl)
	return lyrics
}

// scrapePageForLyrics does the actual work of scraping the page for the lyrics
func scrapePageForLyrics(link string) string {
	var lyrics string
	log.WithField("link", link).Info("webpage::scrapePageForLyrics")
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	err := chromedp.Run(ctx,
		chromedp.Navigate(link),
		chromedp.InnerHTML(".lyrics", &lyrics, chromedp.ByQuery),
	)
	if err != nil {
		log.Error(err)
	}

	return strings.TrimSpace(lyrics)
}
