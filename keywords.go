package lioengine

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"

	r "gopkg.in/dancannon/gorethink.v2"
)

// Keyword struct
type Keyword struct {
	ID     string `gorethink:"id"`
	Word   string `gorethink:"word"`
	Points int    `gorethink:"points"`
}

// this is the function that will be called
// on init() and will initialize the
// keywords var.
func fetchKeywords(errs chan error) {
	if rethink != nil {
		res, err := r.Table(rethink.table).Run(rethink.session)
		if err != nil {
			errs <- err
			return
		}
		defer res.Close()

		err = res.All(&keywords)
		if err != nil {
			errs <- err
			return
		}
		errs <- nil
		return
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

			for i, word := range updt.words {
				updt.words[i] = strings.ToLower(word)
			}
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
			var repeatedWords []string
			for _, word := range updt.words {
			WORD: // Label so the O(n3) is not that bad...
				for _, keyword := range keywords {
					if word == keyword.Word || word == projectName { // Check if matches
						if !isARepeatedWord(word, &repeatedWords) { // WARN: O(n3) HERE <--------------------
							if word == keyword.Word {
								points += keyword.Points
							} else {
								points += 3 // It's not a keyword, but the project name
							}
							repeatedWords = append(repeatedWords, word)
						}
						continue WORD
					}
				}
			}
			updt.points = points
		}(&wg, update)
	}
	wg.Wait()
}

func isARepeatedWord(word string, repeatedWords *[]string) bool {
	if len(*repeatedWords) == 0 {
		*repeatedWords = append(*repeatedWords, word)
		return false
	}
	for _, repWord := range *repeatedWords {
		if word == repWord {
			return true
		}
	}
	return false
}

func filterUpdates(updates *[]*Update) (newUpdates []*Update) {
	for _, update := range *updates {
		if update.points >= minPoints {
			newUpdates = append(newUpdates, update)
			log.Println(update.points)
		}
	}
	return
}

// GetKeywordsNumber returns the lenght of keywords
func GetKeywordsNumber() int {
	return len(keywords)
}

// GetKeyword returns a keyword by it's word
func GetKeyword(keyword string) *Keyword {
	for _, key := range keywords {
		if key.Word == keyword {
			return key
		}
	}
	return nil
}

// AddKeyword adds a keyword
func AddKeyword(word string, points int) (*Keyword, error) {
	if key := GetKeyword(word); key != nil {
		return nil, fmt.Errorf("this keyword is already added")
	}
	if rethink != nil {
		newKeyword := &Keyword{
			ID:     strconv.FormatFloat(rand.Float64(), 'f', 50, 64),
			Word:   word,
			Points: points,
		}

		_, err := r.Table(rethink.table).Insert(newKeyword).RunWrite(rethink.session)
		if err != nil {
			return nil, err
		}

		keywords = append(keywords, newKeyword)
		log.Println("added", word, ":", points)
		return newKeyword, nil
	}
	return nil, nil
}

// ModifyKeyword modifies the points for a keyword
func ModifyKeyword(id string, points int) (*Keyword, error) {
	if rethink != nil {
		_, err := r.Table(rethink.table).Get(id).Update(map[string]interface{}{
			"points": points,
		}).RunWrite(rethink.session)
		if err != nil {
			return nil, err
		}

		var word string
		for _, kw := range keywords {
			if kw.ID == id {
				kw.Points = points
				word = kw.Word
				log.Println("modified", word, ":", points)
				return kw, nil
			}
		}
	}
	return nil, nil
}
