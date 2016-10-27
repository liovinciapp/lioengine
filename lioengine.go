// Package lioengine is a bot for finding updates on already existing projects.
package lioengine

import (
	"log"
	"strings"
)

// sets if using appengine
var usingAppengine bool

// supportedProviders contains all the providers supported by this bot.
var supportedProviders = []string{"Bing", "Twitter"}

// keywords will be initialized on
// the init() func.
var keywords []*Keyword

// minPoints is the minimun value for results
// to be considered as updates.
var minPoints = 15

func init() {
	var errs = make(chan error, 1)
	defer close(errs)
	fetchKeywords(errs)
	var err = <-errs
	if err != nil {
		log.Fatal("error while fetching keywords:", err)
	}
}

// Update is the exported struct that contains
// all kind of info about the project.
type Update struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Link          string   `json:"link"`
	DatePublished string   `json:"date_published"`
	Img           *Img     `json:"img"`
	Category      string   `json:"category"`
	Sources       []string `json:"sources"`

	points int
	words  []string
	// Maybe more stuff ...
}

// Img contains info about the img/thumbnail
type Img struct {
	Link   string `json:"link"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// UseAppengine uses appengine requests
func UseAppengine() {
	usingAppengine = true
}

// SetMinPoints sets the minimum points
// needed for an update to be considered
// one. The higher the harder for the fetched
// results to 'become' an update (less results).
// And the lower the easier to get more updates.
func SetMinPoints(points int) {
	minPoints = points
	log.Println("minPoints now is", minPoints)
}

// replaceSpaces replaces spaces of text with char if the text contains
// spaces.
func replaceSpaces(text, char string) (newText string) {
	// Checks if the text contains spaces
	if strings.Contains(text, " ") {
		// If the text contains spaces it replaces them with char
		newText = strings.Replace(text, " ", char, -1)
		return
	}
	// If text doesn't contain spaces, return the same text
	return text
}
