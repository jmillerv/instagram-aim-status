[![Codacy Badge](https://app.codacy.com/project/badge/Grade/0ae9a3a6678447749c95e617a86be622)](https://www.codacy.com/gh/jmillerv/instagram-away-message/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=jmillerv/instagram-away-message&amp;utm_campaign=Badge_Grade)

# instagram-away-mesage
A command line tool for generating instagram note messages from songs

## What does it do? 
It retrieves a random song based on the artist you enter and returns the first 60 characters. 

## How It Works  
This was made in a mad flash in service of a bit I was doing 
on instagram. I wouldn't call this quality software. If this ends up being of any use to anyone, it will have done far more than I anticipated. 

As this is programmed it in an argument for a given artist, then returns a random track from the body of the returned tracks. From there it navigates to the shazam web page and parses the html looking for the lryics.Once it obtains the lyrics it trims them down to the first 60 characters. It will cut off words, so be warned. 

I didn't build in any verification that it ends up returning the artist you intend.

## Usage 

I've tested this on a handful of bands and so far it's been able to work. 

```azure
COMMANDS:
   get-artist, get  get-artist {string}
   help, h          Shows a list of commands or help for one command

```

For example, the `./app get-artist 3oh!3` will yield `bash: !3: event not found`

Escaped you'll get a response with lyrics

### Special Characters

If you get an error from a band with a special character, try escaping it with a `\`

## Required Dependencies 

### Go
If you've never worked with Go before, you'll need to install that, follow their installation guide.
https://go.dev/doc/install

### RapidAPI 
This is using Shazam's API through Rapid API, so you'll need a rapid API account. Apologies.

## Required Environment Variables
If you've never done set environment variables, [a crossplatform guide](https://wisetut.com/how-to-set-environment-variables-for-windows-linux-and-macosx/)

RAPID_API_APP=instagram-aim-message # name of your app 
RAPID_API_REQUEST_URL=rapidapi.com 
RAPID_API_KEY={token} # X-RapidAPI-Key header parameter
RAPID_API_HOST=shazam.core.p.rapidapi.com # X-RapidAPI-Host parameter

## Building
Once you've gotten every thing setup navigate to the `instagram-aim-status/src` and run `go build` 
It will output a file called `src`. I would personally rename this to something you'll remember like `note-gen` 

ensure that it's executable. On linux that is `chmod+x ./{name of file}`

## Running 

Once that is all setup, you should be able to do `./note-gen get {name of band}`
