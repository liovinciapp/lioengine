package lioengine

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
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

// makeApiCall connects to the apiServer, and fetch all the results
// for later ai ml algorithms.
func makeAPICall(projectName string) (err error) {

	var wg sync.WaitGroup
	for _, provider := range currentProviders {
		wg.Add(1)
		go func() {
			if provider.RequestInfo.Request == nil {
				log.Println("Nil request on makeAPICall")
				provider.RequestInfo.Request, err = http.NewRequest("GET", parseURL(provider.RequestInfo.urlWithParameters, provider.RequestInfo.urlParameters), nil)
			}

			client := new(http.Client)
			response, err := client.Do(provider.RequestInfo.Request)
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

			if provider.Result == nil {
				log.Println("Nil Result b4 unmarshal")
				provider.Result = make(map[string]interface{})
			}

			if err = json.Unmarshal(data, &provider.Result); err != nil {
				log.Printf("Error ocurred at requests.go - json.Unmarshal(...) : %s", err.Error())
				return
			}
			
			log.Println(provider.Result)

			wg.Done()
		}()
	}
	wg.Wait()
	return
}

// parseURL puts the parameters on the url
func parseURL(urlWithParameters string, parameters []string) (parsedURL string) {

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
