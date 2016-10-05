package lioengine

import (
	"errors"
	"log"
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
	Result map[interface{}]interface{}
}

// providers contains all the providers supported by this bot.
var providers []string

// currentProvider for news.
var currentProvider provider

// SetProvider sets the news provider by the name given and
// initializes it with the corresponding apiToken.
// This function should be called before FindUpdatesFor().
// Currently supported providers are: Bing.
func SetProvider(newProviderName, apiToken string) (err error) {
	var oldProvider = currentProvider
	for _, providerName := range providers {
		if providerName == newProviderName {
			getProviderByName(providerName, apiToken)
		}
	}
	if oldProvider.Name == currentProvider.Name {
		err = errors.New("The given provider is not supported.")
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error ocurred at providers.go - SetProvider(...) ")
			log.Println(err)
		}
	}()

	return
}

// getProviderByName returns a provider with the given name.
func getProviderByName(name, apiToken string) {
	switch name {
	case "Bing":
		newBingProvider(apiToken)
		break
	default:
		// You should never ever get here.
		// For any reason.
		// Like, NEVER, even if the provider name is worng.
		break
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error ocurred at providers.go - getProviderByName(...) ")
			log.Println(err)
		}
	}()

	return
}

// setupProviders adds the supported providers to the
// supported providers list.
func setupProviders() {
	providers = append(providers, "Bing")
}
