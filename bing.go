package lioengine

import (
	"net/http"
)

// Adds bing provider support for the bot.
type Bing struct {
	provider
}

// Returns a new initialized bing provider.
func NewBingProvider(apiToken string) (bing *Bing) {
	bing.Name = "Bing"
	bing.Token = apiToken
	bing.setupDefaultRequestInfo()
	return
}

func (p *Bing) setupDefaultRequestInfo() {
	p.RequestInfo.url = "https://api.cognitive.microsoft.com/bing/v5.0/news/search?"
	parameters := []string{"q", "count"}
	p.RequestInfo.urlParameters = parameters
	p.RequestInfo.Request = p.setupDefaultHttpRequest()
	p.RequestInfo.urlWithParameters = p.setupDefaultUrlWithParameters()
	return
}

func (p *Bing) setupDefaultHttpRequest() (request *http.Request) {
	request.Header.Add("Ocp-Apim-Subscription-Key", p.Token)
	return
}

func (p *Bing) setupDefaultUrlWithParameters() {

}
