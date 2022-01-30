package hn_og_combo

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	hn "github.com/hughmandeville/hnui/pkg/hackernews"
	"github.com/otiai10/opengraph/v2"
)

type Item struct {
	ID           int                  `json:"id"`
	Title        string               `json:"title"`
	URL          string               `json:"url"`
	Time         int                  `json:"time"`
	Image        string               `json:"image"`
	Icon         string               `json:"icon"`
	Publisher    string               `json:"publisher"`
	Description  string               `json:"description"`
	By           string               `json:"by"`
	Points       int                  `json:"points"`
	CommentCount int                  `json:"comment_count"`
	HNItem       *hn.Item             `json:"hn_item"`
	OGItem       *opengraph.OpenGraph `json:"og_item"`
}

func GetTopStories(numStories int, verbose bool) (items []Item, err error) {
	// Get top stories from Hacker News.
	hnItems, err := hn.GetTopStories(numStories)
	if err != nil {
		log.Fatalf("Problem getting top stories: %s", err)
		return
	}

	// Get Open Graph data.
	for _, hi := range hnItems {
		item := Item{
			ID:           hi.ID,
			URL:          hi.URL,
			Title:        hi.Title,
			Time:         hi.Time,
			By:           hi.By,
			Points:       hi.Score,
			CommentCount: len(hi.Kids),
			HNItem:       &hi,
		}
		time.Sleep(100 * time.Millisecond)

		// Get Open Graph data.
		// https://pkg.go.dev/github.com/otiai10/opengraph/v2
		// TBD: set timeout.
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		ogi, err := opengraph.Fetch(hi.URL)
		if err == nil {
			item.OGItem = ogi
		}

		item.Description = ogi.Description
		item.Icon = sanitizeURL(item.URL, ogi.Favicon.URL)

		// Set image.
		if len(ogi.Image) > 0 {
			item.Image = sanitizeURL(item.URL, ogi.Image[0].URL)
		}

		// Set publisher.
		item.Publisher = strings.TrimSpace(ogi.SiteName)

		// Fix bad data.
		correctData(&item)
		items = append(items, item)

		if verbose {
			fmt.Printf(" %9d  %-30s  %s\n", item.ID, item.Publisher, item.Title)
		}
	}
	return
}

// Fix known images with icon, image, and publisher data.
func correctData(item *Item) {

	if item.Title == "" {
		item.Title = item.OGItem.Title
	}
	item.Title = strings.TrimSpace(item.Title)

	// Get URL's domain name and remove www.
	domain := ""
	pu, err := url.Parse(item.URL)
	if err == nil {
		domain = strings.TrimPrefix(pu.Hostname(), "www.")
	}

	// if image http remove
	if strings.HasPrefix(item.Image, "http:") {
		item.Image = ""
	}

	// if icon http remove
	if strings.HasPrefix(item.Icon, "http:") {
		item.Icon = ""
	}

	// set icon if missing for some well known publishers
	if item.Icon == "" {
		switch strings.ToLower(domain) {
		case "news.ycombinator.com":
			item.Icon = "https://news.ycombinator.com/favicon.ico"
		case "npr.org":
			item.Icon = "https://www.npr.org/favicon.ico"
		case "ourworldindata.org":
			item.Icon = "https://ourworldindata.org/favicon.ico"
		case "wpr.org":
			item.Icon = "https://www.wpr.org/sites/default/files/favicon_0_0.ico"
		}
	}

	// fix broken icons of some well known publishers
	switch item.Icon {
	case "https://www.bloomberg.com/favicon.ico":
		item.Icon = "https://assets.bwbx.io/s3/javelin/public/hub/images/favicon-black-63fe5249d3.png"
	case "https://news.ycombinator.com/item/favicon.ico":
		item.Icon = "https://news.ycombinator.com/favicon.ico"
	}

	if item.Publisher == "" {
		item.Publisher = domain
	}

	// fix publisher name for some well known publishers
	switch strings.ToLower(item.Publisher) {
	case "bbc.com":
		item.Publisher = "BBC"
	case "bloomberg.com":
		item.Publisher = "Bloomberg"
	case "business-standard.com":
		item.Publisher = "Business Standard"
	case "developer.apple.com":
		item.Publisher = "Apple Developer"
	case "ge.com":
		item.Publisher = "GE"
	case "hudsonreview.com":
		item.Publisher = "The Hudson Review"
	case "kaggle.com":
		item.Publisher = "Kaggle"
	case "nasdaq.com":
		item.Publisher = "Nasdaq"
	case "nature.com":
		item.Publisher = "Nature"
	case "news.ycombinator.com":
		item.Publisher = "Hacker News"
	case "nytimes.com":
		item.Publisher = "The New York Times"
	case "thelocal.com":
		item.Publisher = "The Local"
	case "vice.com":
		item.Publisher = "Vice"
	}

	// Shorten long publisher names with pipe symbol (|) by removing text after pipe symbol.
	i := strings.Index(item.Publisher, "|")
	if len(item.Publisher) > 20 && i > 3 {
		item.Publisher = strings.TrimSpace(item.Publisher[:i-1])
	}

	if len(item.Publisher) > 32 {
		item.Publisher = item.Publisher[:29] + "…"
	}

	// crop description at 300 characters
	if len(item.Description) > 300 {
		item.Description = item.Description[:297] + "…"
	}

	// unset bad descriptions
	if item.Description == "We can’t find the page you are looking for." {
		item.Description = ""
	}

	item.Description = strings.TrimSpace(item.Description)
}

// Turn relative URLs into absolute URLs (/foo/bar.jpg -> https://example.com/foo/bar.jpg).
func sanitizeURL(parentURL string, childURL string) (sanitizedURL string) {
	sanitizedURL = strings.TrimSpace(childURL)
	if sanitizedURL == "" || strings.HasPrefix(sanitizedURL, "http:") || strings.HasPrefix(sanitizedURL, "https:") {
		return
	}
	if strings.HasPrefix(childURL, "//") {
		sanitizedURL = fmt.Sprintf("https:%s", childURL)
		return
	}
	pu, err := url.Parse(parentURL)
	if err != nil {
		return
	}
	if strings.HasPrefix(childURL, "/") {
		sanitizedURL = fmt.Sprintf("%s://%s%s", pu.Scheme, pu.Hostname(), childURL)
		return
	}
	path := pu.Path
	pi := strings.LastIndex(path, "/")
	if pi > 0 {
		path = path[:pi]
	}
	sanitizedURL = fmt.Sprintf("%s://%s%s/%s", pu.Scheme, pu.Hostname(), path, childURL)
	return
}
