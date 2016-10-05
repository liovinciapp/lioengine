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

// apiToken is the private token for api authentication.
const apiToken = "TOKEN"

// apiServer is the request url path we have to connect.
const apiServer = "https://api.cognitive.microsoft.com/bing/v5.0/news/search?q=%s"

// apiRequest sets up the request to the apiServer
var apiRequest = &http.Request{}

// Project is the exported struct that contains
// all kind of info about the project.
type Project struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Img         Img       `json:"img"`
	// Maybe more stuff ...
}

// result is the struct that matches results overflow from
// api calls.
type result struct {
	name string
	url  string
	img  Img
}

type Img struct {
	Link string
}

// FindUpdatesFor is the 'main' and probably the only exported
// function this package will have.
func FindUpdatesFor(projectName string) (project *Project, err error) {
	return
}

// makeApiCall connects to the apiServer, and fetch all the results
// for later ai ml algorithms.
func makeApiCall(projectName string) (results []*result, httpStatus int, err error) {
	url := fmt.Sprintf(apiServer, projectName)
	response, err := apiClient
	return
}

// replaceSpaces replaces spaces by a '+' symbol.
func replaceSpaces(text string) (fixedText string) {
	if strings.Contains(text, " ") {
		fixedText = strings.Replace(text, " ", "+", 0)
		return
	}
	return text
}
