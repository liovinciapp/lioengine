package lioengine

import (
	"sync"

	"github.com/Shixzie/bingnews"
)

// Adds bing provider support for the bot.
// this is just a wraping struct.
type bingProv struct {
	// Name is the name of the news provider.
	Name string
	// Token used to authenticate to the api.
	Token string
	// Engine is the bingnews engine used to request the
	// API.
	Engine *bingnews.Engine
	// Config is the config for the Engine.
	Config *bingnews.Config
	// Results will contain every result fetched from
	// the API.
	Results []*bingnews.Result
}

// newProvider creates a ready to use bing provider.
func (b *bingProv) newProvider(apiToken string, count int) (err error) {

	// For bing we'll use github.com/Shixzie/bingnews, so
	// b.RequestInfo.Request will be nil because we don't need it,
	// also b.urlKeys will be empty, again, thanks to the imported
	// bingnews package.

	b.Name = "Bing"
	b.Token = apiToken
	b.Config = bingnews.NewConfig(apiToken)
	b.Config.Count = count
	b.Config.Mkt = "en-US"
	b.Config.Appengine = usingAppengine
	b.Engine, err = bingnews.NewEngine(b.Config)
	if err != nil {
		return
	}
	return
}

// search calls to the provider api and fetch results into
// b.Result
func (b *bingProv) search(projectName string, wg *sync.WaitGroup, errs chan error) {
	defer wg.Done()
	var err error
	b.Results, err = b.Engine.SearchFor(projectName)
	if err != nil {
		errs <- err
		return
	}
	errs <- nil
}

// standarize converts the fetched results from the provider
// into a []*Update
func (b *bingProv) standarize(updates *[]*Update, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, result := range b.Results {
		newUpdate := &Update{}
		newUpdate.Title = result.Title
		newUpdate.Description = result.Description
		newUpdate.Link = result.Link
		newUpdate.DatePublished = result.PubDate
		if result.Img != nil {
			newUpdate.Img = new(Img)
			newUpdate.Img.Link = result.Img.ContentURL
			newUpdate.Img.Width = result.Img.Width
			newUpdate.Img.Height = result.Img.Height
		} else {
			newUpdate.Img = new(Img)
		}
		newUpdate.Category = result.Category
		newUpdate.Sources = result.Sources
		*updates = append(*updates, newUpdate)
	}
	return
}
