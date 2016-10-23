[![Wercker](https://img.shields.io/wercker/ci/wercker/docs.svg?maxAge=2592000)]()  [![GoDoc](https://godoc.org/github.com/Shixzie/lioengine?status.svg)](https://godoc.org/github.com/Shixzie/lioengine)

# lioengine
liovinci's bot for finding updates on already existing projects.

## Status
Bot should be ready for use. 

## Goals
1. Automate the multi-bot search proccess so users don't have to make it.
2. Add support for more storage engines, such as mysql, postgresql, and mongo. 
3. Some stuff that i can't remember right now.

## How it works
Basically, the bot fetches *n* number of results from the supported providers and puts the results into one *common struct* for text analysis.

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

	lioengine.SetDatabase("rethinkdb",
	 []string{
		 "localhost", // the hosts
	 },
	 "test", 		  // database name
	 "keywords", 	  // table name
	)

	var numbersOfTwitterResults, numbersOfBingResults int
	numbersOfTwitterResults = 3
	numbersOfBingResults = 5
	
	// Sets Bing as our news/updates provider for bots 1, 2 and 3.
	lioengine.AddUpdatesProvider("Bing", "API TOKEN", numbersOfBingResults, bot1, bot2, bot3 ...)

	// Can add more than 1, but right now Bing is the only supported provider.
	// For twitter we need the OAuth2 Token, cuz the bot uses application-only as it doesn't need
	// user behavior.
	lioengine.AddUpdatesProvider("Twitter", "OAuth2 Token", numbersOfTwitterResults, bot1, bot2, bot3)

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
	// ...
}
```