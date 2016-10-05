package lioengine

import (
	"./bing"
	"errors"
)

// Provider is the struct that is used to work with
// different news providers.
type Provider struct {
	// Name is the name of the news provider.
	Name string
	// RequestInfo contains everything related to the resquest made
	// to the api.
	RequestInfo *ApiRequest
}

// providers contains all the providers supported by this bot.
var providers []*Provider

// currentProvider for news.
var currentProvider *Provider

// SetProvider updates the currentProvider to the new one.
func SetProvider(newProviderName string) (err error) {
	var oldProvider = currentProvider
	for _, provider := range providers {
		if provider.Name == newProviderName {
			currentProvider = provider
		}
	}
	if oldProvider == currentProvider {
		err = errors.New("The provider given is not supported.")
		return
	}
	return
}

func setupProviders() {
	providers = append(providers, bing.Setup())
}
