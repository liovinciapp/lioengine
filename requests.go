package lioengine

import (
	"net/http"
	"strings"
	"strconv"
	"reflect"
)

// apiRequest contains all info about the request
// that will be executed to the api.
type apiRequest struct {
	// urlKeys used on the request.
	urlKeys []string //eg. []string {"key1", "key2"}
	// Request is used to set extra info, such as headers.
	Request *http.Request
	// Quantity is the number of results to be fetch
	Quantity quantity
}

type quantity int

// String returns q value as a string
func (q quantity) String() string {
	ref := reflect.ValueOf(q)
	return strconv.Itoa(int(ref.Int()))
}

// Int returns q value as an int
func (q quantity) Int() int {
	ref := reflect.ValueOf(q)
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
