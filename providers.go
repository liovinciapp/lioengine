package lioengine

import "errors"

// provider is the struct that is used to work with
// different news providers.
type provider struct {
	// Name is the name of the news provider.
	Name string
	// Token used to authenticate to the api.
	Token string
	// RequestInfo contains everything related to the resquest made
	// to the api.
	RequestInfo apiRequest
	// Result is a placeholder for the api results.
	Result map[string]interface{}
}

// providers contains all the providers supported by this bot.
var supportedProviders = []string{"Bing"}

// currentProviders for news.
var currentProviders = []*provider{}

// AddUpdatesProvider adds the news provider by the name given and
// initializes it with the corresponding apiToken.
// This function should be called before FindUpdatesFor().
// Currently supported providers are: Bing.
func AddUpdatesProvider(newProviderName, apiToken string) (err error) {
	var alreadyAdded = false
	// Iterates through all currentProviders
	for _, currentProvider := range currentProviders {
		// Checks if we already have this provider
		if newProviderName == currentProvider.Name {
			alreadyAdded = true
		}
	}

	// If we do, then return
	if alreadyAdded {
		err = errors.New("This provider is already added.")
		return
	}

	var itsASupportedProvider = false
	// Iterates through all our supported providers
	for _, supportedProviderName := range supportedProviders {
		if newProviderName == supportedProviderName {
			itsASupportedProvider = true
		}
	}

	// If is one of our supported providers and we haven't added it yet, then we add it.
	if itsASupportedProvider {
		setupProvider(newProviderName, apiToken)
	} else { // If provider not supported.
		err = errors.New("This provider is not supported by the bot.")
		return
	}

	return
}


// setupProvider generates a provider corresponding to it's name
func setupProvider(providerName, apiToken string) {
	switch providerName {
	case "Bing":
		bing := bing{}
		provider := bing.setup(apiToken)
		currentProviders = append(currentProviders, &provider)
		return
	}
	return
}
