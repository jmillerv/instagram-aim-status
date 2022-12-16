package parser

// GetFirst60Characters takes in the shazamUrl and returns the first 60 characters of lyrics
func GetFirst60Characters(shazamUrl string) string {
	lyrics := getLyricsFromWebPage(shazamUrl)
	return trimLyrics(lyrics)
}

func trimLyrics(lyrics string) string {
	return lyrics[:60]
}
