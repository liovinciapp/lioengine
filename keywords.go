package lioengine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

// Keyword struct
type Keyword struct {
	ID     int    `json:"id"`
	Word   string `json:"word"`
	Points int    `json:"points"`
}

// this is the function that will be called
// on init() and will initialize the
// keywords var.
func fetchKeywords(errs chan error) {
	data, err := ioutil.ReadFile("keywords.json")
	if err != nil {
		errs <- err
		return
	}
	err = json.Unmarshal(data, &keywords)
	if err != nil {
		errs <- err
		return
	}
	log.Println("len of keywords:", len(keywords))
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
	wg.Add(len(*updates))
	for _, update := range *updates { // Iterate through all updates
		go func(waitg *sync.WaitGroup, updt *Update) { // Make it go faster
			defer waitg.Done() // Notice our progress senpai
			var points int     // How many we have
			var repeatedWords []string
		WORD: // Label so the O(n3) is not that bad...
			for _, word := range updt.words {
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

// GetKeywords returns all current keywords
func GetKeywords() []*Keyword {
	return keywords
}

// GetKeywordsCount returns the lenght of keywords
func GetKeywordsCount() int {
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

// GetKeywordByID returns a keyword by it's id
func GetKeywordByID(id int) *Keyword {
	for _, key := range keywords {
		if key.ID == id {
			return key
		}
	}
	return nil
}

var lastID = 700

// AddKeyword adds a keyword
func AddKeyword(word string, points int) (*Keyword, error) {
	if key := GetKeyword(word); key != nil {
		return nil, fmt.Errorf("this keyword is already added")
	}
	newKeyword := &Keyword{
		ID:     lastID,
		Word:   word,
		Points: points,
	}
	keywords = append(keywords, newKeyword)
	lastID++
	return newKeyword, nil
}

// ModifyKeyword modifies the points for a keyword
func ModifyKeyword(id int, points int) (*Keyword, error) {

	var word string
	key := GetKeywordByID(id)
	if key == nil {
		return nil, fmt.Errorf("the keyword wasn't found")
	}
	key.Points = points
	word = key.Word
	log.Println("modified", word, ":", points)
	return key, nil

}
