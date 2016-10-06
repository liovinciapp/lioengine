// Package lioengine is a ml bot that will find updates for the
// project name you give it.
package lioengine

import (
	"time"
)

func init() {
}

// Update is the exported struct that contains
// all kind of info about the project.
type Update struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	URL           string    `json:"url"`
	DatePublished time.Time `json:"date_published"`
	Img           Img       `json:"img"`
	Category      string    `json:"category"`
	// Maybe more stuff ...
}

// Img contains info about the img/thumbnail
type Img struct {
	Link   string `json:"link"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}


