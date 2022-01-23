package main

import (
	"encoding/json"
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

const (
	NumStories = 10
	OutFile    = "hn_topstories.json"
	Verbose    = true
)

type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Icon        string `json:"icon"`
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Kids        []int  `json:"kids"`
	Publisher   string `json:"publisher"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// Get top 70 Hacker News stories. If there are no errors, writes to tn_topstories.json.
// Calls the Hacker News API.
//   https://github.com/HackerNews/API
// Uses a Go library to get additional Open Graph data for the article (image, icon, and publisher).
//   https://github.com/otiai10/opengraph
// To Do:
//   - Support command line flags to set out file and verbosity.
func main() {
	ids, err := getTopStories()
	if err != nil {
		log.Fatalf("Problem getting top stories: %s", err)
		return
	}
	var items []Item
	for i, id := range ids {
		time.Sleep(200 * time.Millisecond)
		if i >= NumStories {
			break
		}
		item, err := getItem(id)
		if err != nil {
			log.Fatalf("Problem getting item: %s", err)
			return
		}
		addOGData(&item)
		if Verbose {
			fmt.Printf("%9d  %s\n", item.ID, item.Title)
		}
		items = append(items, item)
	}
	data, err := json.Marshal(items)
	if err != nil {
		log.Fatalf("Problem marshalling items: %s", err)
		return
	}

	err = os.WriteFile(OutFile, data, 0644)
	if err != nil {
		log.Fatalf("Problem saving to file: %s", err)
		return
	}
}

// Add Open Graph data to the item (image, icon, and publisher).
func addOGData(item *Item) (err error) {
	ogp, err := opengraph.Fetch(item.URL)
	if err != nil {
		return
	}
	item.Icon = sanitizeURL(item.URL, ogp.Favicon.URL)
	if len(ogp.Image) > 0 {
		item.Image = sanitizeURL(item.URL, ogp.Image[0].URL)
	}
	item.Publisher = ogp.SiteName
	return
}

// Turn relative URLs into absolute URLs (/foo/bar.jpg -> https://example.com/foo/bar.jpg).
func sanitizeURL(parentURL string, childURL string) (sanitizedURL string) {
	sanitizedURL = childURL
	if childURL == "" || strings.HasPrefix(childURL, "http:") || strings.HasPrefix(childURL, "https:") {
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
	client := &http.Client{Timeout: 1 * time.Second}
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
	return
}

// Get top stories from the Hacker News API.
func getTopStories() (itemIDs []int, err error) {
	tsAPI := "https://hacker-news.firebaseio.com/v0/topstories.json"
	req, err := http.NewRequest("GET", tsAPI, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: 1 * time.Second}
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
