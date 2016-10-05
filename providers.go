package lioengine

import (
	"errors"
	"github.com/Shixzie/bingprovider"
)

// Provider is the struct that is used to work with
// different news providers.
type Provider struct {
	// Name is the name of the news provider.
	Name string
	// Token used to authenticate to the api.
	Token string
	// RequestInfo contains everything related to the resquest made
	// to the api.
	RequestInfo *ApiRequest
	// Result is a placeholder for the api results.
	Result map[string]interface{}
}

// providers contains all the providers supported by this bot.
var providers []string

// currentProvider for news.
var currentProvider *Provider

// SetProvider updates the currentProvider to the new one.
func SetProvider(newProviderName, apiToken string) (err error) {
	var oldProvider = currentProvider
	for _, providerName := range providers {
		if providerName == newProviderName {
			currentProvider = getProviderByName(providerName, apiToken)
		}
	}
	if oldProvider == currentProvider {
		err = errors.New("The given provider is not supported.")
		return
	}
	return
}

// SetProviderToken sets the token for the provider specified with
// SetProvider(). This needs to be called
func SetProviderToken(token string) {
	currentProvider.RequestInfo.Token = token
}

// getProviderByName returns a provider with the given name.
func getProviderByName(name, apiToken string) (provider *Provider) {
	switch name {
	case "Bing":
		provider = NewBingProvider(apiToken)
		break
	}
	return
}

// setupProviders adds the supported providers to the
// supported providers list.
func setupProviders() {
	providers = append(providers, "Bing")
}
