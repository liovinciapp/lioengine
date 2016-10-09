package lioengine

import (
	"bytes"
	"net/http"
	"net/url"
	"sync"
	"io/ioutil"
	"log"
	"encoding/json"
)

// Adds bing provider support for the bot.
// this is just a wraping struct.
type bingProv struct{}

// setup returns a new initialized bing provider.
func (b bingProv) setup(apiToken string, count int) (prov *provider) {
	prov = b.newProvider(apiToken, count)
	return
}

// newProvider creates a ready to use bing provider.
func (b bingProv) newProvider(apiToken string, count int) (prov *provider) {
	prov = &provider{}
	prov.Name = "Bing"
	prov.Token = apiToken
	prov.Result = make(map[string]interface{})
	prov.Type = &bingProv{}

	// Sets a non nil value to RequestInfo
	prov.RequestInfo = new(apiRequest)

	// Sets a non nil value to RequestInfo.Request, the http method defaults to GET
	prov.RequestInfo.Request, _ = http.NewRequest("", "https://api.cognitive.microsoft.com/bing/v5.0/news/search", nil)

	keys := []string{ // Maybe in the future the user could choose what keys use.
		"q",
		"count",
		"mkt",
	} 
	prov.RequestInfo.urlKeys = keys
	prov.RequestInfo.Count = countType(count) // Number of results to be fetched
	b.setupDefaultHTTPRequest(apiToken, prov.RequestInfo.Request)

	return
}

// setupDefaultHttpRequest customizes the http request in order to
// have a successful call to the api.
func (b bingProv) setupDefaultHTTPRequest(apiToken string, req *http.Request) {
	// We set the Ocp-Apim-Subscription-Key needed to authenticate to the api.
	req.Header.Add("Ocp-Apim-Subscription-Key", apiToken)

	// Sets protocol
	req.Proto = "HTTP/1.1"
	// Sets scheme
	req.URL.Scheme = "https"
	// Sets host
	req.Host = "api.cognitive.microsoft.com"
	return
}

// setupDefaultUrlWithParameters generates the url with parameters to be used
// when makeApiCall() calls to the api.
func (b bingProv) addParamsToURL(urlKeys []string, urlValues []string, url *url.URL) {
	// We use a buffer because it's the most efficient way to
	// concatenate strings
	var buffer bytes.Buffer
	var isLastIteration = false
	for index, key := range urlKeys {
		if index == len(urlKeys)-1 {
			isLastIteration = true
		}
		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(urlValues[index])
		if !isLastIteration {
			buffer.WriteString("&")
		}
	}
	url.RawQuery = buffer.String()
}

// search calls to the provider api and fetch results into
// prov.Result
func (b bingProv) search(projectName string, prov *provider, wg *sync.WaitGroup) (err error) {
	nonSpacedProjectName := replaceSpaces(projectName, "+") // Replace spaces with the + symbol

	// Have to match the order defined on prov.RequestInfo.urlKeys
	// when setting up the provider.
	urlValues := []string {
		nonSpacedProjectName, // Query -> q
		prov.RequestInfo.Count.String(), // Count -> count
		"en-US", // Mkt -> mkt
	}
	b.addParamsToURL(prov.RequestInfo.urlKeys, urlValues, prov.RequestInfo.Request.URL)
	resp, err := http.DefaultClient.Do(prov.RequestInfo.Request)
	if err != nil {
		wg.Done()
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wg.Done()
		return
	}
	
	err = json.Unmarshal(data, &prov.Result)
	if err != nil {
		wg.Done()
		return
	}
	log.Println(prov.Result)
	wg.Done()
	return
}
