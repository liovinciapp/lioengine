package lioengine

import (
	"bytes"
)

// Adds bing provider support for the bot.
type bing struct {
	p provider
}

// Returns a new initialized bing provider.
func newBingProvider(apiToken string) (bing *bing) {
	bing.p.Name = "Bing"
	bing.p.Token = apiToken
	bing.setupDefaultRequestInfo()
	return
}

// setupDefaultRequestInfo
func (bing *bing) setupDefaultRequestInfo() {
	bing.p.RequestInfo.url = "https://api.cognitive.microsoft.com/bing/v5.0/news/search?"
	parameters := []string{"q", "count"} // Maybe in the future the user could choose what parameters use.
	bing.p.RequestInfo.urlParameters = parameters
	bing.p.RequestInfo.Quantity = 10 // Fetch first 10 results
	bing.setupDefaultUrlWithParameters()

	// This needs to be called the last, so we can use the generated url with parameters.
	bing.setupDefaultHttpRequest()
	return
}

// setupDefaultHttpRequest customizes the http request in order to
// have a successful call to the api.
func (bing *bing) setupDefaultHttpRequest() {
	// We set the Ocp-Apim-Subscription-Key needed to authenticate to the api.
	bing.p.RequestInfo.Request.Header.Add("Ocp-Apim-Subscription-Key", provider.Token)

	// We set the request url so when executing makeApiCall(), we use the right url path.
	bing.p.RequestInfo.Request.URL.Path = provider.RequestInfo.urlWithParameters
}

// setupDefaultUrlWithParameters generates the url with parameters to be used
// when makeApiCall() calls to the api.
func (bing *bing) setupDefaultUrlWithParameters() {
	var buffer bytes.Buffer
	buffer.WriteString(bing.p.RequestInfo.url)
	var isLastIteration = false
	for index, parameter := range bing.p.RequestInfo.urlParameters {
		if index == len(bing.p.RequestInfo.urlParameters)-1 {
			isLastIteration = true
		}
		buffer.WriteString(parameter)
		buffer.WriteString("=")
		buffer.WriteString("%v")
		if !isLastIteration {
			buffer.WriteString("&")
		}
	}
	bing.p.RequestInfo.urlWithParameters = buffer.String()
}
