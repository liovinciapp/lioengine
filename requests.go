package lioengine

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// apiRequest contains all info about the request
// that will be executed to the api.
type apiRequest struct {
	// url we'll make the request on
	url string
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

// makeApiCall connects to the apiServer, and fetch all the results
// for later ai ml algorithms.
func makeApiCall(projectName string) (err error) {
	client := &http.Client{}
	response, err := client.Do(currentProvider.RequestInfo.Request)
	if err != nil {
		log.Printf("Error ocurred at requests.go - client.Do(...) : %s", err.Error())
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error ocurred at requests.go - ioutil.ReadAll(...) : %s", err.Error())
		return
	}
	if err = json.Unmarshal(data, &currentProvider.Result); err != nil {
		log.Printf("Error ocurred at requests.go - json.Unmarshal(...) : %s", err.Error())
		return
	}
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
