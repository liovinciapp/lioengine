package lioengine

import (
    //"github.com/dghubble/go-twitter/twitter"
    //"golang.org/x/oauth2"
    "sync"
)

// Adds twitter provider support for the bot.
// this is just a wraping struct.
type twitterProv struct{}

// setup returns a new initialized twitter provider.
func (t twitterProv) setup(apiToken string) (prov *provider) {
    prov = t.newProvider(apiToken)
    return
}

// newProvider creates a ready to use bing provider.
func (t twitterProv) newProvider(apiToken string) (prov *provider) {

    // For twitter we'll use github.com/dghubble/go-twitter/twitter, so
    // prov.RequestInfo.Request will be nil because we don't need it,
    // also prov.urlKeys will be empty, again, thanks to the imported
    // twitter package.

    prov = &provider{}
    prov.Name = "Twitter"
    prov.Token = apiToken
    prov.Result = make(map[string]interface{})
    prov.Type = twitterProv{}
    // Sets a non nil value to RequestInfo
	prov.RequestInfo = new(apiRequest)
    prov.RequestInfo.Quantity = 10 // Fetch first 10 results
    return
}

// search calls to the provider api and fetch results into
// prov.Result
func (t twitterProv) search(projectName string, prov *provider, wg *sync.WaitGroup) {

}