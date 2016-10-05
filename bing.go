package lioengine

import (
	"bytes"
	"fmt"
	"net/http"
)

// Adds bing provider support for the bot.
type bing struct {
	provider
}

// Returns a new initialized bing provider.
func newBingProvider(apiToken string) (bing *bing) {
	bing.Name = "Bing"
	bing.Token = apiToken
	bing.setupDefaultRequestInfo()
	return
}

// setupDefaultRequestInfo
func (provider *bing) setupDefaultRequestInfo() {
	provider.RequestInfo.url = "https://api.cognitive.microsoft.com/bing/v5.0/news/search?"
	parameters := []string{"q", "count"} // Maybe in the future the user could choose what parameters use.
	provider.RequestInfo.urlParameters = parameters
	provider.RequestInfo.Quantity = 10 // Fetch first 10 results
	provider.setupDefaultUrlWithParameters()

	// This needs to be called the last, so we can use the generated url with parameters.
	provider.setupDefaultHttpRequest()
	return
}

// setupDefaultHttpRequest customizes the http request in order to
// have a successful call to the api.
func (provider *bing) setupDefaultHttpRequest() {
	// We set the Ocp-Apim-Subscription-Key needed to authenticate to the api.
	provider.RequestInfo.Request.Header.Add("Ocp-Apim-Subscription-Key", provider.Token)

	// We set the request url so when executing makeApiCall(), we use the right url path.
	provider.RequestInfo.Request.URL.Path = provider.RequestInfo.urlWithParameters
}

// setupDefaultUrlWithParameters generates the url with parameters to be used
// when makeApiCall() calls to the api.
func (provider *bing) setupDefaultUrlWithParameters() {
	var buffer bytes.Buffer
	buffer.WriteString(provider.RequestInfo.url)
	var isLastIteration = false
	for index, parameter := range provider.RequestInfo.urlParameters {
		if index == len(provider.RequestInfo.urlParameters)-1 {
			isLastIteration = true
		}
		buffer.WriteString(parameter)
		buffer.WriteString("=")
		buffer.WriteString("%v")
		if !isLastIteration {
			buffer.WriteString("&")
		}
	}
	provider.RequestInfo.urlWithParameters = buffer.String()
}
