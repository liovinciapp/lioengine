package lioengine

import (
	"net/http"
)

// ApiRequest contains all info about the request
// that will be executed to the api.
type apiRequest struct {
	// Url we'll make the request on
	url string
	// UrlParameters used on the request.
	urlParameters []string
	// Request is used to set extra info, such as headers.
	Request *http.Request
	// Quantity is the number of result to be fetch
	Quantity int
	// urlWithParameters is the result of successfuly generating
	// the request path.
	urlWithParameters string
}
