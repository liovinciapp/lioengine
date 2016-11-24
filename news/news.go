package news

import (
	"errors"
	"time"

	"github.com/SlyMarbo/rss"
)

// Listener listens for news on the hosts rss
type Listener struct {
	// unreadItems is a map of [feedTitle]item
	// for keeping track of unread items
	unreadItems []*rss.Item
	// readItems is a map of [feedTitle]item too
	// for keeping track of already read items
	readItems []*rss.Item
	// feeds contains all feeds
	feeds []*rss.Feed
}

// NewListener creates a new Listener and returns an error
func NewListener(hosts []string) (*Listener, error) {
	l := new(Listener)
	for _, host := range hosts {
		err := l.AddHost(host)
		if err != nil {
			return nil, err
		}
	}
	return l, nil
}

// AddHost adds a host to the listener
func (l *Listener) AddHost(host string) error {
	for _, f := range l.feeds {
		if f.UpdateURL == host {
			return errors.New("this host already exists")
		}
	}
	feed, err := rss.Fetch(host)
	if err != nil {
		return err
	}
	l.feeds = append(l.feeds, feed)
	return nil
}

func (l *Listener) updateFeed(feed *rss.Feed) error {
	for !time.Now().After(feed.Refresh) {
		time.Sleep(time.Minute * 10)
	}

	err := feed.Update()
	if err != nil {
		l.updateFeed(feed)
	}
	for _, i := range feed.Items {
		l.addItem(i)
	}
	return nil
}

// MarkAsUnread marks the item as unread
func (l *Listener) MarkAsUnread(item *rss.Item) {
	var newSlice []*rss.Item
	for _, i := range l.readItems {
		if i.Title != item.Title && i.Link != item.Link {
			newSlice = append(newSlice, i)
		}
	}
	l.readItems = newSlice
	l.unreadItems = append(l.unreadItems, item)
}

// MarkAsRead marks the item as read
func (l *Listener) MarkAsRead(item *rss.Item) {
	var newSlice []*rss.Item
	for _, i := range l.unreadItems {
		if i.Title != item.Title && i.Link != item.Link {
			newSlice = append(newSlice, i)
		}
	}
	l.unreadItems = newSlice
	l.readItems = append(l.readItems, item)
}

func (l *Listener) addItem(item *rss.Item) {
	var exists bool
	for _, i := range l.unreadItems {
		if exists {
			return
		}
		exists = i.ID == item.ID && i.Title == item.Title && i.Link == item.Link
	}
	l.unreadItems = append(l.unreadItems, item)
}

// GetReadItems returns the read items
func (l *Listener) GetReadItems() []*rss.Item { return l.readItems }

// GetUnreadItems returns the unread items
func (l *Listener) GetUnreadItems() []*rss.Item { return l.unreadItems }

// GetItem returns an item by it's name and ID
func (l *Listener) GetItem(title string, ID string, read bool) (*rss.Item, error) {
	if read {
		for _, i := range l.readItems {
			if i.Title == title && i.ID == ID {
				return i, nil
			}
		}
	} else {
		for _, i := range l.unreadItems {
			if i.Title == title && i.ID == ID {
				return i, nil
			}
		}
	}
	return nil, errors.New("item not found")
}

// deleteOldReadItems deletes all read items that are 10 days older
// or more
func (l *Listener) deleteOldReadItems() {
	for {
		var newSlice []*rss.Item
		for _, i := range l.readItems {
			iTime := i.Date
			if !time.Now().After(iTime.Add(time.Hour * 240)) {
				newSlice = append(newSlice, i)
			}
		}
		l.readItems = newSlice
		time.Sleep(time.Hour * 240)
	}
}

// Listen listens for news
func (l *Listener) Listen() {
	go l.deleteOldReadItems()
	for {
		for _, f := range l.feeds {
			for _, i := range f.Items {
				l.addItem(i)
			}
			go l.updateFeed(f) // fix this, ignoring err
		}
		time.Sleep(time.Minute * 10)
	}
}
