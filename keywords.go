package lioengine

import (
	r "gopkg.in/dancannon/gorethink.v2"
)

// These fileds needs to be 'Exported' 'cuz 
// of the gorethink package
type keyword struct {
	ID     string `gorethink:"id"`
	Word   string `gorethink:"word"`
	Points int    `gorethink:"points"`
}

// keywords will be initialized on
// the init() func.
var keywords []*keyword


// this is the function that will be called
// on init() and will initialize the
// keywords var
func fetchKeywords() (err error) {
	// logic to fetch keywords...
	session, err := r.Connect(r.ConnectOpts{
		// if there's only 1 host we use Address,
		// if not, we use Addresses
		Address:  "localhost",

		// Addresses: []string{
		// 	"host1",
		// 	"host2",
		// 	"host3",
		// },

		// set the database so we don't need
		// to call r.Db('') everytime
		Database: "test", 
	})
	if err != nil {
		return err
	}

	res, err := r.Table("keywords").Run(session)
	if err != nil {
		return err
	}
	defer res.Close()

	err = res.All(&keywords)
	if err != nil {
		return err
	}
	return
}

func (b *Bot) analizeUpdates(projectName string) (updates []*Update, err error) {

	err = splitWords(&b.results)

	calculatePoints(&b.results)

	updates = filterUpdates(&b.results)

	return
}

func splitWords(updates *[]*Update) (err error) {
	// logic ...
	return
}

func calculatePoints(updates *[]*Update) {
	// logic ...
}

func filterUpdates(updates *[]*Update) (newUpdates []*Update) {
	// logic ...
	return
}