// Package lioengine is a bot for finding updates on already existing projects.
package lioengine

import (
	"fmt"
	"strings"
)

// supportedProviders contains all the providers supported by this bot.
var supportedProviders = []string{"Bing", "Twitter"}

// keywords will be initialized on
// the init() func.
var keywords []*keyword

// minPoints is the minimun value for results
// to be considered as updates.
var minPoints = 15

// database is the current database used
// for storing the keywords
var database string

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

// SetMinPoints sets the minimum points
// needed for an update to be considered
// one. The higher the harder for the fetched
// results to 'become' an update (less results).
// And the lower the easier to get more updates.
func SetMinPoints(points int) {
	minPoints = points
}

// SetDatabase sets the database that will
// be used.
//
// name is the storage engine that we'll use.
// valid names are: rethinkdb.
//
// hosts is the host for the given engine.
//
// database is the database name where the keywords
// are located.
//
// table is the table name where the keywords are
// located. The table should have this structure:
// tableName: {
// 	word   string
// 	points int
// }
func SetDatabase(name string, hosts []string, database string, table string) (err error) {
	lowerName := strings.ToLower(name)
	switch lowerName {
	case "rethinkdb":
		database = lowerName
		rethink, err = configRethinkdb(hosts, database, table)
		return err
	default:
		err = fmt.Errorf("%s database is not supported", lowerName)
		return err
	}
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
