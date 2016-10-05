package lioengine

import (
	"bytes"
	"log"
)

// Adds bing provider support for the bot.
type bing struct {
	p provider
}

// Returns a new initialized bing provider.
func newBingProvider(apiToken string) {
	var bing = new(bing)
	bing.p.Name = "Bing"
	bing.p.Token = apiToken
	bing.p.Result = make(map[interface{}]interface{})
	bing.setupDefaultRequestInfo()

	currentProvider = bing.p

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error ocurred at bing.go - dunno ")
			log.Println(err)
		}
	}()
}

// setupDefaultRequestInfo
func (b bing) setupDefaultRequestInfo() (newBing *bing) {
	newBing.p.RequestInfo.url = "https://api.cognitive.microsoft.com/bing/v5.0/news/search?"
	parameters := []string{"q", "count"} // Maybe in the future the user could choose what parameters use.
	newBing.p.RequestInfo.urlParameters = parameters
	newBing.p.RequestInfo.Quantity = 10 // Fetch first 10 results
	b.setupDefaultUrlWithParameters(&newBing)

	// This needs to be called the last, so we can use the generated url with parameters.
	b.setupDefaultHttpRequest()
	return
}

// setupDefaultHttpRequest customizes the http request in order to
// have a successful call to the api.
func (bing bing) setupDefaultHttpRequest() {
	// We set the Ocp-Apim-Subscription-Key needed to authenticate to the api.
	bing.p.RequestInfo.Request.Header.Add("Ocp-Apim-Subscription-Key", bing.p.Token)

	// We set the request url so when executing makeApiCall(), we use the right url path.
	bing.p.RequestInfo.Request.URL.Path = bing.p.RequestInfo.urlWithParameters
}

// setupDefaultUrlWithParameters generates the url with parameters to be used
// when makeApiCall() calls to the api.
func (b bing) setupDefaultUrlWithParameters(newBing *bing) {
	var buffer bytes.Buffer
	buffer.WriteString(newBing.p.RequestInfo.url)
	var isLastIteration = false
	for index, parameter := range newBing.p.RequestInfo.urlParameters {
		if index == len(newBing.p.RequestInfo.urlParameters)-1 {
			isLastIteration = true
		}
		buffer.WriteString(parameter)
		buffer.WriteString("=")
		buffer.WriteString("%v")
		if !isLastIteration {
			buffer.WriteString("&")
		}
	}
	newBing.p.RequestInfo.urlWithParameters = buffer.String()
}
