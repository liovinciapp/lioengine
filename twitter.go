package lioengine

import (
    "github.com/dghubble/go-twitter/twitter"
    "golang.org/x/oauth2"
    "sync"
    "log"
    "net/http"
)

// Adds twitter provider support for the bot.
// this is just a wraping struct.
type twitterProv struct{
    client *twitter.Client
}

// setup returns a new initialized twitter provider.
func (t twitterProv) setup(apiToken string, count int) (prov *provider) {
    prov = t.newProvider(apiToken, count)
    return
}

// newProvider creates a ready to use bing provider.
func (t twitterProv) newProvider(apiToken string, count int) (prov *provider) {

    // For twitter we'll use github.com/dghubble/go-twitter/twitter, so
    // prov.RequestInfo.Request will be nil because we don't need it,
    // also prov.urlKeys will be empty, again, thanks to the imported
    // twitter package.

    prov = &provider{}
    prov.Name = "Twitter"
    prov.Token = apiToken
    prov.Result = make(map[string]interface{})
    prov.Type = &twitterProv{}
    // Sets a non nil value to RequestInfo
	prov.RequestInfo = new(apiRequest)
    prov.RequestInfo.Count = countType(count) /// Number of results to be fetched
    return
}

// search calls to the provider api and fetch results into
// prov.Result
func (t twitterProv) search(projectName string, prov *provider, wg *sync.WaitGroup) (err error) {
    switch v := prov.Type.(type) {
        case *twitterProv:
            // If we haven't already initialized the client, we create one
            if v.client == nil {
                config := &oauth2.Config{}
                token := &oauth2.Token{AccessToken: prov.Token}
                // http.Client will automatically authorize Requests
                httpClient := config.Client(oauth2.NoContext, token)
                v.client = twitter.NewClient(httpClient)
            }
            search, resp, err := v.client.Search.Tweets(&twitter.SearchTweetParams{
                Query: projectName,
                Count: prov.RequestInfo.Count.Int(),
            })
            if err != nil {
                log.Println(err.Error())
                wg.Done()
                return err
            }
            log.Println(search.Statuses)
            if resp.StatusCode == http.StatusOK {
                prov.Result = search.Statuses
            }
            log.Println(prov.Result)
            break
    }
    wg.Done()
    return
}