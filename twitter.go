package lioengine

import (
	"net/http"
	"sync"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
)

// Adds twitter provider support for the bot.
// this is just a wraping struct.
type twitterProv struct {
	client *twitter.Client
}

// setup returns a new initialized twitter provider.
func (t twitterProv) setup(apiToken string, count int) (prov *provider, err error) {
	prov, err = t.newProvider(apiToken, count)
	return
}

// newProvider creates a ready to use bing provider.
func (t twitterProv) newProvider(apiToken string, count int) (prov *provider, err error) {

	// For twitter we'll use github.com/dghubble/go-twitter/twitter.

	prov = &provider{}
	prov.Name = "Twitter"
	prov.Token = apiToken
	prov.Result = []twitter.Tweet{}
	prov.Type = &twitterProv{}
	// Sets a non nil value to RequestInfo
	prov.RequestInfo = new(apiRequest)
	prov.RequestInfo.Count = countType(count) /// Number of results to be fetched

	switch v := prov.Type.(type) {
	case *twitterProv:
		config := &oauth2.Config{}
		token := &oauth2.Token{AccessToken: prov.Token}
		// http.Client will automatically authorize Requests
		httpClient := config.Client(oauth2.NoContext, token)
		v.client = twitter.NewClient(httpClient)
		break
	}

	return
}

// search calls to the provider api and fetch results into
// prov.Result
func (t *twitterProv) search(projectName string, prov *provider, wg *sync.WaitGroup) (err error) {
	switch v := prov.Type.(type) {
	case *twitterProv:
		search, resp, err := v.client.Search.Tweets(&twitter.SearchTweetParams{
			Query: projectName,
			Count: prov.RequestInfo.Count.Int(),
		})
		if err != nil {
			wg.Done()
			return err
		}
		if resp.StatusCode == http.StatusOK {
			prov.Result = search.Statuses
		}
		break
	}
	wg.Done()
	return
}

// standarize converts the fetched results from the provider
// into a []*Update
func (t *twitterProv) standarize(tweets []twitter.Tweet) (updates []*Update) {
	for _, tweet := range tweets {
		newUpdate := &Update{}
		newUpdate.Name = tweet.Source
		newUpdate.Description = tweet.Text
		newUpdate.URL = tweet.Source
		newUpdate.DatePublished = tweet.CreatedAt
		newUpdate.Img = nil
		newUpdate.Category = ""
		newUpdate.Sources = []string{tweet.User.Name}
		newUpdate._type = twitterProv{}
		updates = append(updates, newUpdate)
	}
	return
}
