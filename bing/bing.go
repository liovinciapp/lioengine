package bing

import (
	"../lioengine"
	"net/http"
)

func Setup(apiToken string) (provider *lioengine.Provider) {
	provider.Name = "Bing"
	provider.RequestInfo = setupRequestInfo(apiToken)
}

func setupRequestInfo(apiToken string) (requestInfo *lioengine.ApiRequest) {
	requestInfo.Token = apiToken
	requestInfo.Url = "https://api.cognitive.microsoft.com/bing/v5.0/news/search?"
	parameters := []string{"q", "count"}
	requestInfo.UrlParameters = parameters
	requestInfo.Request = setupHttpRequest(requestInfo.Token)
	return
}

func setupHttpRequest(token string) (request *http.Request) {
	request.Header.Add("Ocp-Apim-Subscription-Key", token)
	return
}

// result is the struct that matches results overflow from
// api calls.
type Result struct {
	Name        string
	Url         string
	Thumbnail   resultThumbnail
	Description string
	Provider    resultProvider
}

type resultThumbnail struct {
	img
}

type resultProvider struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type img struct {
	ContentUrl string `json:"content_url"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}
