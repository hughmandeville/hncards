// This package contains functions for getting the top stories data from Hacker News.
// It uses the Hacker News API.
//   https://github.com/HackerNews/API

package hackernews

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Item data. Has fields from Hacker News top stories API and additional fields from Open Graph.
type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// Get top stories from Hacker News
func GetTopStories(numStories int) (items []Item, err error) {
	itemIDs, err := getTopStoriesIds()
	if err != nil {
		return
	}
	for i, id := range itemIDs {
		if i >= numStories {
			break
		}
		time.Sleep(100 * time.Millisecond)
		item, err := getItem(id)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return
}

// Get top stories item IDs from the Hacker News API.
func getTopStoriesIds() (itemIDs []int, err error) {
	tsAPI := "https://hacker-news.firebaseio.com/v0/topstories.json"
	req, err := http.NewRequest("GET", tsAPI, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	// If error sleep and try again.
	if err != nil {
		time.Sleep(1 * time.Second)
		resp, err = client.Do(req)
		if err != nil {
			return
		}
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

// Get item info from the Hacker News API.
func getItem(id int) (item Item, err error) {
	itemAPI := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	req, err := http.NewRequest("GET", itemAPI, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)

	// If error sleep and try again.
	if err != nil {
		time.Sleep(1 * time.Second)
		resp, err = client.Do(req)
		if err != nil {
			return
		}
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP error %s", resp.Status)
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &item)
	if err != nil {
		return
	}
	// Items with no URL are Hacker News links.
	if item.URL == "" {
		item.URL = fmt.Sprintf("https://news.ycombinator.com/item?id=%d", item.ID)
	}
	return
}
