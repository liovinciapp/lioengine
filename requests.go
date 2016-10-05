package lioengine

import (
	"net/http"
)

// ApiRequest contains all info about the request
// that will be executed to the api.
type ApiRequest struct {
	// Token used to authenticate to the api.
	Token string
	// Url we'll make the request on
	Url string
	// UrlParameters used on the request.
	UrlParameters []string
	// Request is used to set extra info, such as headers.
	Request *http.Request
	// Quantity is the number of result to be fetch
	Quantity int
}
