package lioengine

import (
	"net/http"
	"sync"

	"errors"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
)

// Adds twitter provider support for the bot.
// this is just a wraping struct.
type twitterProv struct {
	// Name is the name of the news provider.
	Name string
	// Token used to authenticate to the api.
	Token string
	// Count is the number of tweets to be fetched
	Count int
	// Client is the twitter client used to request the
	// API.
	Client *twitter.Client
	// Results will contain every result fetched from
	// the API.
	Results []twitter.Tweet
	// httpClient is the http client used by our twitter
	// Client
	httpClient *http.Client
}

// newProvider creates a ready to use twitter provider.
func (t *twitterProv) newProvider(apiToken string, count int) (err error) {

	// For twitter we'll use github.com/dghubble/go-twitter/twitter.

	t.Name = "Twitter"
	t.Token = apiToken

	t.Count = count /// Number of results to be fetched

	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: t.Token}
	// http.Client will automatically authorize Requests
	t.httpClient = config.Client(oauth2.NoContext, token)
	t.Client = twitter.NewClient(t.httpClient)

	return
}

// search calls to the provider api and fetch results into
// prov.Result
func (t *twitterProv) search(projectName string, wg *sync.WaitGroup, errs chan error) {
	var searchParams = new(twitter.SearchTweetParams)
	searchParams.Query = projectName
	searchParams.Count = t.Count
	if t.Client.Search == nil {
		errs <- errors.New("twitter client search is nil")
		wg.Done()
		return
	}
	search, resp, err := t.Client.Search.Tweets(searchParams)
	if err != nil {
		errs <- err
		wg.Done()
		return
	}
	if resp.StatusCode == http.StatusOK {
		t.Results = search.Statuses
	}

	wg.Done()
	return
}

// standarize converts the fetched results from the provider
// into a []*Update
func (t *twitterProv) standarize() (updates []*Update) {
	for _, tweet := range t.Results {
		newUpdate := &Update{}
		newUpdate.Title = ""
		newUpdate.Description = tweet.Text
		newUpdate.Link = tweet.Source
		newUpdate.DatePublished = tweet.CreatedAt
		newUpdate.Img = nil
		newUpdate.Category = ""
		newUpdate.Sources = []string{tweet.User.Name}
		newUpdate._type = &twitterProv{}
		updates = append(updates, newUpdate)
	}
	return
}
