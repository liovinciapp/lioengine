package lioengine

import (
	"errors"
	"sync"

	"github.com/Shixzie/bingnews"
	"github.com/dghubble/go-twitter/twitter"
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
	results          []*Update
}

// FindUpdatesFor is the 'main' function for this package. And will return
// an []*Update for the given project name.
func (b *Bot) FindUpdatesFor(projectName string) (updates []*Update, err error) {
	err = b.makeAPICalls(projectName)
	if err != nil {
		err = errors.New("Failure while executing API calls")
		return
	}
	b.standarizeResults()
	updates = b.results
	//updates, err = b.analizeUpdates()
	// if err != nil {
	// 	err = errors.New("Failure while analysing updates")
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
		switch provType := provider.Type.(type) {
		case *bingProv:
			wg.Add(1)
			go provType.search(projectName, provider, wg)
			break
		case *twitterProv:
			wg.Add(1)
			go provType.search(projectName, provider, wg)
			break
		}
	}
	wg.Wait()
	return
}

// setupProvider generates a provider corresponding to it's name
func (b *Bot) setupProvider(providerName, apiToken string, count int) (err error) {
	switch providerName {
	case "Bing":
		bing := bingProv{}
		provider, err := bing.setup(apiToken, count)
		if err != nil {
			return err
		}
		b.currentProviders = append(b.currentProviders, provider)
		break
	case "Twitter":
		twitter := twitterProv{}
		provider, err := twitter.setup(apiToken, count)
		if err != nil {
			return err
		}
		b.currentProviders = append(b.currentProviders, provider)
		break
	}
	return
}

// Returns the provider by it's name.
func (b *Bot) getProviderByName(providerName string) *provider {
	for _, provider := range b.currentProviders {
		if provider.Name == providerName {
			return provider
		}
	}
	return nil // Return nil if the provider wasn't found
}

func (b *Bot) standarizeResults() {
	// Iterates through all providers
	for _, provider := range b.currentProviders {
		switch provType := provider.Type.(type) {
		// Check if is a twitter provider
		case *twitterProv:
			// Check if the provider results are tweets
			switch tweets := provider.Result.(type) {
			case []twitter.Tweet:
				// Standarizes all fetched tweets
				standarizedTweets := provType.standarize(tweets)
				// Iterates through all standarized tweets
				for _, standarizedTweet := range standarizedTweets {
					// Adds the standarized data to the *Bot.results
					b.results = append(b.results, standarizedTweet)
				}
				break
			}
			break
		// Check if is a bing provider
		case *bingProv:
			// Check if the provider results are bing results
			switch results := provider.Result.(type) {
			case []*bingnews.Result:
				// Standarizes all fetched results
				standarizedResults := provType.standarize(results)
				// Iterates through all standarized results
				for _, standarizedResult := range standarizedResults {
					// Adds the standarized data to the *Bot.results
					b.results = append(b.results, standarizedResult)
				}
				break
			}
			break
		}
	}
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
			err = bot.setupProvider(newProviderName, apiToken, count)
			if err != nil {
				return
			}
		} else { // If provider not supported.
			err = errors.New("This provider is not supported by the bot.")
			return
		}
	}
	return
}
