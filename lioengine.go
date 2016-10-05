// Package lioengine is a ml bot that will find updates for the
// project name you give it.
package lioengine

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func init() {
	// Creates the defaults for all
	// supported providers
	setupProviders()
}

// Update is the exported struct that contains
// all kind of info about the project.
type Update struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Url           string    `json:"url"`
	DatePublished time.Time `json:"date_published"`
	Img           Img       `json:"img"`
	Category      string    `json:"category"`
	Source        source    `json:"source"`
	// Maybe more stuff ...
}

type Img struct {
	Link   string `json:"link"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// FindUpdatesFor is the 'main' function for this package. And will return
// an []*Update for the given project name.
func FindUpdatesFor(projectName string) (updates *Update, err error) {
	return
}

// makeApiCall connects to the apiServer, and fetch all the results
// for later ai ml algorithms.
func makeApiCall(projectName string) (results []*result, httpStatus int, err error) {
	return
}

// replaceSpaces replaces spaces of text with char if the text contains
// spaces.
func replaceSpaces(text, char string) (newText string) {
	if strings.Contains(text, " ") {
		newText = strings.Replace(text, " ", char, 0)
		return
	}
	return text
}
