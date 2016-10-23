package lioengine

import (
	"errors"
	"sync"
)

// NewBot creates a bot.
func NewBot() (bot *Bot) {
	bot = new(Bot)
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
// an []*Update for the given project name and an err.
func (b *Bot) FindUpdatesFor(projectName string) (updates []*Update, err error) {
	err = b.makeAPICalls(projectName)
	if err != nil {
		err = errors.New("failure while executing API calls")
		return nil, err
	}

	b.standarizeResults()

	updates = b.analizeUpdates(projectName)

	return
}

// makeApiCalls connects to the apiServer, and fetch all the results
// for later analysis.
func (b *Bot) makeAPICalls(projectName string) (err error) {

	b.results = []*Update{} // for every search made, reset the results

	// wg waits for all concurrent searches to finish
	// and blocks this func until all of the searches are done.
	var wg sync.WaitGroup
	var errs = make(chan error, 2)
	defer close(errs) // close the chan

	for _, provider := range b.currentProvidersNames {
		switch provider {
		case "Bing":
			// Search with bing
			wg.Add(1)
			go b.bing.search(projectName, &wg, errs)
			break
		case "Twitter":
			// Search with twitter
			wg.Add(1)
			go b.twitter.search(projectName, &wg, errs)
			break
		}
	}

	// check for errs
	for i := 0; i < len(b.currentProvidersNames); i++ {
		err = <-errs
		if err != nil {
			return err
		}
	}

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
	var wg sync.WaitGroup

	wg.Add(2)

	// Adds twitter results
	go b.twitter.standarize(&b.results, &wg)

	// Adds bing results
	go b.bing.standarize(&b.results, &wg)

	wg.Wait()
}

// AddUpdatesProvider adds the news provider by the name given and
// initializes it with the corresponding apiToken.
// This function should be called before FindUpdatesFor().
// This is also designed to be called multiple times.
// Current supported providers are: Bing, Twitter.
//
// For the Twitter provider you'll need an OAuth2 Token.&
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
