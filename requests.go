package lioengine

import (
	"fmt"
	"net/http"
	"strings"
)

// apiRequest contains all info about the request
// that will be executed to the api.
type apiRequest struct {
	// url we'll make the request on
	host string
	path string
	// urlParameters used on the request.
	urlParameters []string
	// Request is used to set extra info, such as headers.
	Request *http.Request
	// Quantity is the number of result to be fetch
	Quantity int
	// urlWithParameters will be the result of successfuly generating
	// the request path with parameters.
	urlWithParameters string
}

// parseURL puts the parameters on the url
func parseURL(urlWithParameters, projectName string, count int) (parsedURL string) {
	noSpacedProjectName := replaceSpaces(projectName, "+")
	parsedURL = fmt.Sprintf(urlWithParameters, noSpacedProjectName, count)
	return
}

// replaceSpaces replaces spaces of text with char if the text contains
// spaces.
func replaceSpaces(text, char string) (newText string) {
	if strings.Contains(text, " ") {
		newText = strings.Replace(text, " ", char, -1)
		return
	}
	return text
}
