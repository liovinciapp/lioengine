package news

import (
	"fmt"
	"testing"
)

func Test_NewListener(t *testing.T) {
	t.Run("listen", testListen)
}

func testListen(t *testing.T) {
	hosts := []string{
		"http://feeds.feedburner.com/TechCrunch/startups",
		"http://feeds.feedburner.com/TechCrunch/social",
		"http://feeds.feedburner.com/Mobilecrunch",
		"http://feeds.feedburner.com/crunchgear",
		"http://feeds.feedburner.com/TechCrunch/Amazon",
		"http://feeds.feedburner.com/TechCrunch/Android",
		"http://feeds.feedburner.com/TechCrunch/Apple",
		"http://feeds.feedburner.com/TechCrunch/Facebook",
		"http://feeds.feedburner.com/TechCrunch/Google",
		"http://feeds.feedburner.com/TechCrunch/Microsoft",
		"http://feeds.feedburner.com/TechCrunch/Samsung",
		"http://feeds.feedburner.com/TechCrunch/Twitter",
		"https://www.wired.com/category/gear/feed/",
		"https://www.wired.com/category/reviews/feed/",
		"https://www.wired.com/category/design/feed/",
		"http://www.techradar.com/rss/news/gaming",
		"http://www.techradar.com/rss/news/portable-devices",
		"http://www.techradar.com/rss/news/phone-and-communications",
		"http://www.techradar.com/rss/reviews/gadgets",
	}
	listener, err := NewListener(hosts)
	if err != nil {
		t.Error(err)
	}
	go listener.Listen()
	fmt.Printf("\n\n")
	for _, i := range listener.GetUnreadItems() {
		fmt.Printf("Title: %s\t\t", i.Title)
		fmt.Printf("Link: %s\n\n", i.Link)
	}
}
