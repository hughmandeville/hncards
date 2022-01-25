package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/otiai10/opengraph/v2"
)

var (
	numStories int
	outFile    string
)

type Item struct {
	By            string `json:"by"`
	Descendants   int    `json:"descendants"`
	Icon          string `json:"icon"`
	ID            int    `json:"id"`
	Image         string `json:"image"`
	Kids          []int  `json:"kids"`
	OGDescription string `json:"og_description"`
	OGTitle       string `json:"og_title"`
	Publisher     string `json:"publisher"`
	Score         int    `json:"score"`
	Time          int    `json:"time"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	URL           string `json:"url"`
}

// Get top 70 Hacker News stories. If there are no errors, writes to tn_topstories.json.
// Calls the Hacker News API.
//   https://github.com/HackerNews/API
// Uses a Go library to get additional Open Graph data for the article (image, icon, and publisher).
//   https://github.com/otiai10/opengraph
// To Do:
//   - Support using the previous file as a cache.
func main() {
	// Parse command line flags.
	flag.IntVar(&numStories, "num", 70, "number of top stories to get")
	flag.StringVar(&outFile, "out", "hn_topstories.json", "output file JSON")
	flag.Parse()

	fmt.Printf("Get Hacker News Top Stories\n")
	fmt.Printf("---------------------------\n")
	fmt.Printf("Out File:    %s\n", outFile)
	fmt.Printf("Num Stories: %d\n\n", numStories)

	ids, err := getTopStories()
	if err != nil {
		log.Fatalf("Problem getting top stories: %s", err)
		return
	}
	var items []Item
	for i, id := range ids {
		time.Sleep(200 * time.Millisecond)
		if i >= numStories {
			break
		}
		item, err := getItem(id)
		if err != nil {
			log.Fatalf("Problem getting item: %s", err)
			return
		}
		addOGData(&item)
		fmt.Printf(" %9d  %s\n", item.ID, item.Title)
		items = append(items, item)
	}
	if len(items) < 10 {
		fmt.Printf("Hacker News API returned less than 10 stories, so not writing to %s.\n", outFile)
		return
	}
	data, err := json.Marshal(items)
	if err != nil {
		log.Fatalf("Problem marshalling items: %s", err)
		return
	}

	err = os.WriteFile(outFile, data, 0644)
	if err != nil {
		log.Fatalf("Problem saving to file: %s", err)
		return
	}
}

// Add Open Graph data to the item (image, icon, and publisher).
func addOGData(item *Item) (err error) {
	// set publisher to the URL's domain name by default
	if item.Publisher == "" {
		pu, err := url.Parse(item.URL)
		if err == nil {
			item.Publisher = pu.Hostname()
			item.Publisher = strings.TrimPrefix(item.Publisher, "www.")
		}
	}

	ogp, err := opengraph.Fetch(item.URL)
	if err != nil {
		return
	}
	icon := sanitizeURL(item.URL, ogp.Favicon.URL)
	if icon != "" {
		if icon == "https://news.ycombinator.com/item/favicon.ico" {
			icon = "https://news.ycombinator.com/favicon.ico"
		}
		item.Icon = icon
	}
	if len(ogp.Image) > 0 {
		item.Image = sanitizeURL(item.URL, ogp.Image[0].URL)
	}
	publisher := strings.TrimSpace(ogp.SiteName)
	if publisher != "" {
		item.Publisher = publisher
	}

	item.OGDescription = ogp.Description
	item.OGTitle = ogp.Title
	return
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

// Get item info from the Hacker News API.
func getItem(id int) (item Item, err error) {
	itemAPI := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	req, err := http.NewRequest("GET", itemAPI, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: 4 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP error %s", resp.Status)
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(b, &item)
	if err != nil {
		return
	}
	if item.URL == "" {
		item.Publisher = "Hacker News"
		item.Icon = "https://news.ycombinator.com/favicon.ico"
		item.URL = fmt.Sprintf("https://news.ycombinator.com/item?id=%d", item.ID)
	}
	return
}

// Get top stories from the Hacker News API.
func getTopStories() (itemIDs []int, err error) {
	tsAPI := "https://hacker-news.firebaseio.com/v0/topstories.json"
	req, err := http.NewRequest("GET", tsAPI, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: 4 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP error %s", resp.Status)
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(b, &itemIDs)
	return
}
