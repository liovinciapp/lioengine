package lioengine

import (
	"bytes"
	"net/http"
	"net/url"
	"sync"
)

// Adds bing provider support for the bot.
type bing struct{}

// Returns a new initialized bing provider.
func (b bing) setup(apiToken string) (prov provider) {
	prov = b.setupDefaultRequestInfo(apiToken)
	return
}

// setupDefaultRequestInfo
func (b bing) setupDefaultRequestInfo(apiToken string) (prov provider) {

	prov.Name = "Bing"
	prov.Token = apiToken
	prov.Result = make(map[string]interface{})
	prov.Type = bing{}

	// Sets a non nil value to RequestInfo
	prov.RequestInfo = new(apiRequest)

	// Sets a non nil value to RequestInfo.Request
	prov.RequestInfo.Request, _ = http.NewRequest("GET", "", nil)

	parameters := []string{"q", "count"} // Maybe in the future the user could choose what parameters use.
	prov.RequestInfo.urlParameters = parameters
	prov.RequestInfo.host = "https://api.cognitive.microsoft.com"
	prov.RequestInfo.path = "/bing/v5.0/news/search"
	prov.RequestInfo.Quantity = 10 // Fetch first 10 results
	prov.RequestInfo.urlWithParameters = b.setupDefaultURLWithParameters(prov.RequestInfo.host, prov.RequestInfo.path, prov.RequestInfo.urlParameters)
	b.setupDefaultHTTPRequest(apiToken, prov.RequestInfo.host, prov.RequestInfo.path, prov.RequestInfo.Request)

	return
}

// setupDefaultHttpRequest customizes the http request in order to
// have a successful call to the api.
func (b bing) setupDefaultHTTPRequest(apiToken, host, path string, req *http.Request) {
	// We set the Ocp-Apim-Subscription-Key needed to authenticate to the api.
	req.Header.Add("Ocp-Apim-Subscription-Key", apiToken)

	// Sets a non nil value to req.URL
	req.URL = new(url.URL)

	// We set the request url so when executing makeApiCall(), we use the right url path.
	req.URL.Host = host
	req.URL.Path = path

	return
}

// setupDefaultUrlWithParameters generates the url with parameters to be used
// when makeApiCall() calls to the api.
func (b bing) setupDefaultURLWithParameters(host, path string, urlParameters []string) (urlWithParameters string) {
	var buffer bytes.Buffer
	buffer.WriteString(host)
	buffer.WriteString(path)
	var isLastIteration = false
	for index, parameter := range urlParameters {
		if index == len(urlParameters)-1 {
			isLastIteration = true
		}
		if index == 0 {
			buffer.WriteString("?")
		}
		buffer.WriteString(parameter)
		buffer.WriteString("=")
		buffer.WriteString("%v")
		if !isLastIteration {
			buffer.WriteString("&")
		}
	}
	urlWithParameters = buffer.String()
	return
}

// search calls to the provider api and fetch results
func (b bing) search(prov *provider, wg *sync.WaitGroup) {
	wg.Done()
}
