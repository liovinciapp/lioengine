package lioengine

import (
	"errors"
	"fmt"
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
var supportedProviders = []string{"Bing"}

// currentProviders for news.
var currentProviders = make(map[string]provider)

// AddUpdatesProvider adds the news provider by the name given and
// initializes it with the corresponding apiToken.
// This function should be called before FindUpdatesFor().
// Currently supported providers are: Bing.
func AddUpdatesProvider(newProviderName, apiToken string) (err error) {
	fmt.Println("Enter AddUpdatesProvider")
	fmt.Println(supportedProviders, "< supportedProviders")
	fmt.Println(currentProviders, "< currentProviders before adding")
	fmt.Println(newProviderName, "< the newProviderName")
	// Iterates through all our supported providers
	for _, supportedProviderName := range supportedProviders {
		// Iterates through all currentProviders
		var alreadyAdded = false
		for currentProviderName := range currentProviders {
			if newProviderName == currentProviderName {
				alreadyAdded = true
			}
		}
		// If is one of our supported providers and we haven't added it yet, then we added.
		if newProviderName == supportedProviderName && !alreadyAdded {
			fmt.Println("Gonna add", newProviderName)
			setupProvider(supportedProviderName, apiToken)
		} else if supportedProviderName != newProviderName { // If provider not supported.
			err = errors.New("This provider is not supported by the bot.")
			return
		} else { // If provider already added.
			err = errors.New("This provider is already added.")
			return
		}
	}
	defer fmt.Println(currentProviders, "< currentProviders after adding")
	return
}

func setupProvider(providerName, apiToken string) {
	switch providerName {
	case "Bing":
		bing := &bing{}
		provider := provider{}
		bing.setup(apiToken, provider)
		currentProviders[providerName] = provider
		fmt.Println("Added provider")
		return
	}
	return
}
