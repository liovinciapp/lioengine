package lioengine

import (
	"bytes"
	"net/http"
)

// Adds bing provider support for the bot.
type bing struct {
	provider
}

// Returns a new initialized bing provider.
func (bing bing) setup (apiToken string, provider provider) {
	provider.Name = "Bing"
	provider.Token = apiToken
	provider.Result = make(map[string]interface{})
	provider.RequestInfo = bing.setupDefaultRequestInfo(provider)
	provider.RequestInfo.Request = bing.setupDefaultHTTPRequest(provider)
	provider.RequestInfo.urlWithParameters = bing.setupDefaultURLWithParameters(provider.RequestInfo)
	return
}


// setupDefaultRequestInfo
func (bing *bing) setupDefaultRequestInfo(prov provider) (reqInfo *apiRequest) {
	parameters := []string{"q", "count"} // Maybe in the future the user could choose what parameters use.
	reqInfo.urlParameters = parameters
	reqInfo.host = "api.cognitive.microsoft.com"
	reqInfo.path = "/bing/v5.0/news/search"
	reqInfo.Quantity = 10 // Fetch first 10 results
	
	return
}

// setupDefaultHttpRequest customizes the http request in order to
// have a successful call to the api.
func (bing *bing) setupDefaultHTTPRequest(prov provider) (req *http.Request) {
	
	// We set the Ocp-Apim-Subscription-Key needed to authenticate to the api.
	req.Header.Add("Ocp-Apim-Subscription-Key", prov.Token)

	// We set the request url so when executing makeApiCall(), we use the right url path.
	
	req.URL.Host = prov.RequestInfo.host
	req.URL.Path = prov.RequestInfo.path
	return
}

// setupDefaultUrlWithParameters generates the url with parameters to be used
// when makeApiCall() calls to the api.
func (bing *bing) setupDefaultURLWithParameters(reqInfo *apiRequest) (urlWithParameters string) {
	var buffer bytes.Buffer
	buffer.WriteString(reqInfo.host)
	buffer.WriteString(reqInfo.path)
	var isLastIteration = false
	for index, parameter := range reqInfo.urlParameters {
		if index == len(reqInfo.urlParameters)-1 {
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
