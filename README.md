[![Wercker](https://img.shields.io/wercker/ci/wercker/docs.svg?maxAge=2592000)]()  [![AUR](https://img.shields.io/aur/license/yaourt.svg?maxAge=2592000)]() [![GoDoc](https://godoc.org/github.com/Shixzie/lioengine?status.svg)](https://godoc.org/github.com/Shixzie/lioengine)

# lioengine
liovinci's machine learning ai that finds updates for projects. This still on a very early stage of development so don't spect to have it working soon.

## Status
Bot base for fetching data form any kind of news/updates provider is ready. I'm starting to work on the Ai related stuff.

## About Ai - Ml
The Ai stuff will be built with a bunch of layers. So lets say, i've a layer for recognizing if the data contains a link, and what that layer does is check if the data contains "http://" (pretty bad). In the future i might want to replace that shitty conditional for a regular expression, so i only have to replace that layer. Every 'action' that the machine performs for analysing data should have it's own layer.

TLDR: **Still on development. Can be built and executed, but wont work.**


## Supported providers

1. Bing
2. Twitter

## Future providers

1. Facebook
2. Google (This wasn't the first one because it's news API is deprecated)
3. Maybe some famous pages, but that wont happen soon.

## Installation
```
go get github.com/Shixzie/lioengine
```

## Usage
```go
package main

import (
	"github.com/Shixzie/lioengine"
	"bufio"
	"os"
	"fmt"
)

func main() {

	// Create bots, as many as you want
	bot1 := lioengine.NewBot()
	bot2 := lioengine.NewBot()
	bot3 := lioengine.NewBot()

	var numbersOfTwitterResultsToBeFetched, numbersOfBingResultsToBeFetched int
	numbersOfTwitterResultsToBeFetched = 3
	numbersOfBingResultsToBeFetched = 5
	
	// Sets Bing as our news/updates provider for bots 1, 2 and 3.
	lioengine.AddUpdatesProvider("Bing", "API TOKEN", numbersOfBingResultsToBeFetched, bot1, bot2, bot3 ...)

	// Can add more than 1, but right now Bing is the only supported provider.
	// For twitter we need the OAuth2 Token, cuz the bot uses application-only as it doesn't need
	// user behavior.
	lioengine.AddUpdatesProvider("Twitter", "OAuth2 Token", numbersOfTwitterResultsToBeFetched, bot1, bot2, bot3)

	// Creates a reader so we can read from the console.
	// Instead of using the os.Stdin we could use a JSON request to get
	// the project name or get the value from a html form.
	reader := bufio.NewReader(os.Stdin)

	// We ask for the project name we want to get updates for.
    fmt.Print("Enter project name: ")

    // Wait until the user press Enter.
    project, _ := reader.ReadString('\n')
	project = project[:len(project)-2] // Removes \n from input

    // Search for updates.
	// This makes all 3 bots find news for the same project
	// the idea is that you make the bots search for different projects
	// concurrently, so you don't have 2 bots waiting while 1 is searching
	// and analysing 
	updates1, _ := bot1.FindUpdatesFor(project)
	// You can call FindUpdatesFor() ad many times you want for each bot.
	// The only thing is that you need to wait the bot to return updates. 
	updates2, _ := bot2.FindUpdatesFor(project)
	updates3, _ := bot3.FindUpdatesFor(project)

	// Iterates through updates
	for _, update := range updates1 {
		// do whatever with the update info
	}

	for _, update := range updates2 {
		// ...
	}

	for _, update := range updates3 {
		// ...
	}
	...
}
```

## Documentation
####    [Godoc](http://godoc.org/github.com/Shixzie/lioengine)