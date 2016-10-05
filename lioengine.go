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

// Project is the exported struct that contains
// all kind of info about the project.
type Project struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Url           string    `json:"url"`
	DatePublished time.Time `json:"date_published"`
	Img           img       `json:"img"`
	Category      string    `json:"category"`
	Source        source    `json:"source"`
	// Maybe more stuff ...
}

type img struct {
	Link   string `json:"link"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// FindUpdatesFor is the 'main' function this package will have.
func FindUpdatesFor(projectName string) (project *Project, err error) {
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
