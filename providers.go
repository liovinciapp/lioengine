package lioengine

// provider is the struct that is used to work with
// different news providers.
type provider struct {
	// Name is the name of the news provider.
	Name string
	// Token used to authenticate to the api.
	Token string
	// RequestInfo contains everything related to the resquest made
	// to the api.
	RequestInfo *apiRequest
	// Result is a placeholder for the api results.
	Result interface{}
	// Defines if its Bing, Twitter...
	Type interface{}
}

// providers contains all the providers supported by this bot.
var supportedProviders = []string{"Bing", "Twitter"}

