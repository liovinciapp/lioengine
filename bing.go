package lioengine

import (
	"sync"

	"github.com/Shixzie/bingnews"
)

// Adds bing provider support for the bot.
// this is just a wraping struct.
type bingProv struct {
	engine *bingnews.Engine
}

// setup returns a new initialized bing provider.
func (b bingProv) setup(apiToken string, count int) (prov *provider, err error) {
	prov, err = b.newProvider(apiToken, count)
	return
}

// newProvider creates a ready to use bing provider.
func (b bingProv) newProvider(apiToken string, count int) (prov *provider, err error) {

	// For bing we'll use github.com/Shixzie/bingnews, so
	// prov.RequestInfo.Request will be nil because we don't need it,
	// also prov.urlKeys will be empty, again, thanks to the imported
	// bingnews package.

	prov = &provider{}
	prov.Name = "Bing"
	prov.Token = apiToken
	prov.Result = []*bingnews.Result{}
	prov.Type = &bingProv{}

	config := bingnews.NewConfig(apiToken)
	config.Count = count
	config.Mkt = "en-US"

	switch v := prov.Type.(type) {
	case *bingProv:
		v.engine, err = bingnews.NewEngine(config)
		if err != nil {
			return
		}
		break
	}
	return
}

// search calls to the provider api and fetch results into
// prov.Result
func (b *bingProv) search(projectName string, prov *provider, wg *sync.WaitGroup) (err error) {
	switch provType := prov.Type.(type) {
	case *bingProv:
		prov.Result, err = provType.engine.SearchFor(projectName)
		break
	}
	wg.Done()
	return
}

// standarize converts the fetched results from the provider
// into a []*Update
func (b *bingProv) standarize(results []*bingnews.Result) (updates []*Update) {
	for _, result := range results {
		newUpdate := &Update{}
		newUpdate.Name = result.Title
		newUpdate.Description = result.Description
		newUpdate.URL = result.Link
		newUpdate.DatePublished = result.PubDate
		newUpdate.Img = new(Img)
		newUpdate.Img.Link = result.Img.ContentURL
		newUpdate.Img.Width = result.Img.Width
		newUpdate.Img.Height = result.Img.Height
		newUpdate.Category = result.Category
		newUpdate.Sources = result.Sources
		newUpdate._type = bingProv{}
		updates = append(updates, newUpdate)
	}
	return
}
