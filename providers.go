package lioengine

import (
	"errors"
)

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
	Result map[string]interface{}
}

// providers contains all the providers supported by this bot.
var providers []string

// currentProvider for news.
var currentProvider *provider

// SetProvider sets the news provider by the name given and
// initializes it with the corresponding apiToken.
// This function should be called before FindUpdatesFor().
// Currently supported providers are: Bing.
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

// getProviderByName returns a provider with the given name.
func getProviderByName(name, apiToken string) (provider *provider) {
	switch name {
	case "Bing":
		provider = newBingProvider(apiToken)
		break
	default:
		// You should never ever get here.
		// For any reason.
		// Like, NEVER, even if the provider name is worng.
		break
	}
	return
}

// setupProviders adds the supported providers to the
// supported providers list.
func setupProviders() {
	providers = append(providers, "Bing")
}
