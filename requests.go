package lioengine

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// apiRequest contains all info about the request
// that will be executed to the api.
type apiRequest struct {
	// urlKeys used on the request.
	urlKeys []string //eg. []string {"key1", "key2"}
	// Request is used to set extra info, such as headers.
	Request *http.Request
	// Count is the number of results to be fetch
	Count countType
}

type countType int

// String returns c value as a string
func (c countType) String() string {
	ref := reflect.ValueOf(c)
	return strconv.Itoa(int(ref.Int()))
}

// Int returns c value as an int
func (c countType) Int() int {
	ref := reflect.ValueOf(c)
	return int(ref.Int())
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
