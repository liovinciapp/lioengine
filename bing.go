package lioengine

import (
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

func (p *bing) setupDefaultRequestInfo() {
	p.RequestInfo.url = "https://api.cognitive.microsoft.com/bing/v5.0/news/search?"
	parameters := []string{"q", "count"}
	p.RequestInfo.urlParameters = parameters
	p.RequestInfo.Request = p.setupDefaultHttpRequest()
	p.RequestInfo.urlWithParameters = p.setupDefaultUrlWithParameters()
	return
}

func (p *bing) setupDefaultHttpRequest() (request *http.Request) {
	request.Header.Add("Ocp-Apim-Subscription-Key", p.Token)
	return
}

func (p *bing) setupDefaultUrlWithParameters() {

}
