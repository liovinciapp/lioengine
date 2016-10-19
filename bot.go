package lioengine

import (
	"errors"
	"sync"
)

// NewBot creates a bot.
func NewBot() (bot *Bot) {
	bot = new(Bot)
	bot.currentProvidersNames = []string{}
	bot.bing = new(bingProv)
	bot.twitter = new(twitterProv)
	return
}

// Bot holds all utility for searching news and updates
type Bot struct {
	currentProvidersNames []string
	results               []*Update
	bing                  *bingProv
	twitter               *twitterProv
}

// FindUpdatesFor is the 'main' function for this package. And will return
// an []*Update for the given project name.
func (b *Bot) FindUpdatesFor(projectName string) (updates []*Update, err error) {
	err = b.makeAPICalls(projectName)
	if err != nil {
		err = errors.New("failure while executing API calls")
		return
	}
	b.standarizeResults()
	updates = b.results
	// err = b.analizeUpdates(projectName)
	// if err != nil {
	// 	err = errors.New("Failure while analysing updates")
	// 	return
	// }
	return
}

// makeApiCalls connects to the apiServer, and fetch all the results
// for later analysis.
func (b *Bot) makeAPICalls(projectName string) (err error) {

	b.results = []*Update{} // for every search made, reset the results

	// wg waits for all concurrent searches to finish
	// and blocks this func until all of the searches are done.
	var wg sync.WaitGroup
	wg.Add(len(b.currentProvidersNames))
	var errs = make(chan error, len(b.currentProvidersNames))
	defer close(errs)
	wg.Wait()
	return
}

// setupProvider generates a provider corresponding to it's name
func (b *Bot) setupProvider(providerName, apiToken string, count int) (err error) {
	switch providerName {
	case "Bing":
		err = b.bing.newProvider(apiToken, count)
		if err != nil {
			return err
		}
		b.currentProvidersNames = append(b.currentProvidersNames, "Bing")
		break
	case "Twitter":
		err = b.twitter.newProvider(apiToken, count)
		if err != nil {
			return err
		}
		b.currentProvidersNames = append(b.currentProvidersNames, "Twitter")
		break
	}
	return
}

func (b *Bot) standarizeResults() {
	// Adds twitter results
	// Standarizes all fetched tweets
	standarizedTweets := b.twitter.standarize()
	// Iterates through all standarized tweets
	for _, standarizedTweet := range standarizedTweets {
		// Adds the standarized data to the *Bot.results
		b.results = append(b.results, standarizedTweet)
	}

	// #####################################

	// Adds bing results
	// Standarizes all fetched results
	standarizedResults := b.bing.standarize()
	// Iterates through all standarized results
	for _, standarizedResult := range standarizedResults {
		// Adds the standarized data to the *Bot.results
		b.results = append(b.results, standarizedResult)
	}
}

// AddUpdatesProvider adds the news provider by the name given and
// initializes it with the corresponding apiToken.
// This function should be called before FindUpdatesFor().
// This is also designed to be called multiple times.
// Current supported providers are: Bing, Twitter.
//
// For the Twitter provider you'll need an OAuth2 Token.
// Because this bot uses the Application-only mode instead
// of Application-user, check https://dev.twitter.com/oauth
// for more info.
func AddUpdatesProvider(newProviderName, apiToken string, count int, bots ...*Bot) (err error) {
	//Iterates through all bots to add the providers
	for _, bot := range bots {
		var alreadyAdded = false
		// Iterates through all currentProvidersNames
		for _, currentProvider := range bot.currentProvidersNames {
			// Checks if we already have this provider
			if newProviderName == currentProvider {
				alreadyAdded = true
			}
		}

		// If we do, then return
		if alreadyAdded {
			err = errors.New("this provider is already added")
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
			err = errors.New("this provider is not supported by the bot")
			return
		}
	}
	return
}
