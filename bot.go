package lioengine

import (
	"errors"
	"log"
	"sync"
)

// NewBot creates a bot.
func NewBot() (bot *Bot) {
	bot = new(Bot)
	bot.currentProviders = []*provider{}
	return
}

// Bot holds all utility for searching news and updates
type Bot struct {
	currentProviders []*provider
}

// FindUpdatesFor is the 'main' function for this package. And will return
// an []*Update for the given project name.
func (b *Bot) FindUpdatesFor(projectName string) (updates []*Update, err error) {
	err = b.makeAPICalls(projectName)
	if err != nil {
		log.Printf("Error ocurred at lioengine.go - makeApiCalls(...) : %s", err.Error())
		return
	}
	//b.standarizeResults()
	//updates, err = b.analizeUpdates()
	// if err != nil {
	// 	log.Printf("Error ocurred at lioengine.go - analizeUpdates(...) : %s", err.Error())
	// 	return
	// }
	return
}

// makeApiCalls connects to the apiServer, and fetch all the results
// for later ai ml algorithms.
func (b *Bot) makeAPICalls(projectName string) (err error) {
	// wg waits for all concurrent searches to finish
	// and blocks this func until all of the searches are done.
	var wg = new(sync.WaitGroup)
	for _, provider := range b.currentProviders {
		switch v := provider.Type.(type) {
		case *bingProv:
			log.Println("Searching", projectName, "with bing.")
			wg.Add(1)
			go v.search(projectName, provider, wg)
			break
		case *twitterProv:
			log.Println("Searching", projectName, "with twitter.")
			wg.Add(1)
			go v.search(projectName, provider, wg)
			break 
		}
	}
	wg.Wait()
	return
}

// setupProvider generates a provider corresponding to it's name
func (b *Bot) setupProvider(providerName, apiToken string, count int) {
	switch providerName {
	case "Bing":
		bing := bingProv{}
		provider := bing.setup(apiToken, count)
		b.currentProviders = append(b.currentProviders, provider)
		break
	case "Twitter":
		twitter := twitterProv{}
		provider := twitter.setup(apiToken, count)
		b.currentProviders = append(b.currentProviders, provider)
		break
	}
	return
}

// Returns the slide index for the provider with the name
// providerName.
func (b *Bot) getProviderIndexByName(providerName string) (index int) {
	for index, currentProvider := range b.currentProviders {
		if currentProvider.Name == providerName {
			return index
		}
	}
	return
}

// AddUpdatesProvider adds the news provider by the name given and
// initializes it with the corresponding apiToken.
// This function should be called before FindUpdatesFor().
// This is also designed to be called multiple times.
// Current supported providers are: Bing.
func AddUpdatesProvider(newProviderName, apiToken string, count int, bots ...*Bot) (err error) {
	//Iterates through all bots to add the providers
	for _, bot := range bots {
		var alreadyAdded = false
		// Iterates through all currentProviders
		for _, currentProvider := range bot.currentProviders {
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
			bot.setupProvider(newProviderName, apiToken, count)
		} else { // If provider not supported.
			err = errors.New("This provider is not supported by the bot.")
			return
		}
	}
	return
}
