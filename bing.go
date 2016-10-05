package lioengine

import (
	"net/http"
)

// Adds bing provider support for the bot.
type Bing struct {
	Provider
}

func NewBingProvider(apiToken string) (bing *Bing) {
	bing.Token = apiToken
	return
}

func (p *Bing) Setup() (provider Provider) {
	p.Name = "Bing"
	p.RequestInfo = p.setupDefaultRequestInfo()
	return
}

func (p *Bing) setupDefaultRequestInfo() (requestInfo ApiRequest) {
	requestInfo.url = "https://api.cognitive.microsoft.com/bing/v5.0/news/search?"
	parameters := []string{"q", "count"}
	requestInfo.urlParameters = parameters
	requestInfo.Request = setupDefaultHttpRequest()
	return
}

func (p *Bing) setupDefaultHttpRequest() (request *http.Request) {
	request.Header.Add("Ocp-Apim-Subscription-Key", p.Token)
	return
}
