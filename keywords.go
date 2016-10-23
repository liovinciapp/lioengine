package lioengine

import (
	"strings"
	"sync"

	r "gopkg.in/dancannon/gorethink.v2"
)

// These fileds needs to be 'Exported' 'cuz
// of the gorethink package
type keyword struct {
	ID     string `gorethink:"id"`
	Word   string `gorethink:"word"`
	Points int    `gorethink:"points"`
}

// this is the function that will be called
// on init() and will initialize the
// keywords var.
func fetchKeywords(errs chan error) {
	switch database {
	case "rethinkdb":
		res, err := r.Table(rethink.table).Run(rethink.session)
		if err != nil {
			errs <- err
		}
		defer res.Close()

		err = res.All(&keywords)
		if err != nil {
			errs <- err
		}
		break
	}
	errs <- nil
}

func (b *Bot) analizeUpdates(projectName string) (updates []*Update) {

	splitWords(&b.results)

	calculatePoints(&b.results, projectName)

	updates = filterUpdates(&b.results)

	return
}

func splitWords(updates *[]*Update) {
	var wg sync.WaitGroup
	for _, update := range *updates { // Iterate through all updates
		wg.Add(1)
		go func(waitg *sync.WaitGroup, updt *Update) { // Split words concurrently
			defer waitg.Done()                                                   // Let waitg know when we're done
			var allWords []string                                                // just allWords
			allWords = append(allWords, strings.Split(updt.Description, " ")...) // Add description words
			allWords = append(allWords, strings.Split(updt.Title, " ")...)       // Add title words
			allWords = append(allWords, updt.Sources...)                         // Add the sources

			updt.words = append(updt.words, allWords...) // Adds all words to the words field
		}(&wg, update)
	}
	wg.Wait()
	return
}

func calculatePoints(updates *[]*Update, projectName string) {
	var wg sync.WaitGroup
	for _, update := range *updates { // Iterate through all updates
		wg.Add(1)
		go func(waitg *sync.WaitGroup, updt *Update) { // Make it go faster
			defer waitg.Done() // Notice our progress senpai
			var points int     // How many we have
		WORD: // Label so the O(n2) is not that bad...
			for _, word := range updt.words { // WARN: O(n2) HERE <--------------------
				for _, keyword := range keywords { // WARN: O(n2) HERE <--------------------
					if word == keyword.Word || word == projectName { // Check if matches
						points += keyword.Points // Sum points
						continue WORD
					}
				}
			}
			updt.points = points
		}(&wg, update)
	}
	wg.Wait()
}

func filterUpdates(updates *[]*Update) (newUpdates []*Update) {
	for _, update := range *updates {
		if update.points >= minPoints {
			newUpdates = append(newUpdates, update)
		}
	}
	return
}
