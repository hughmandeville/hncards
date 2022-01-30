package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hughmandeville/hnui/pkg/hn_og_combo"
)

var (
	numStories int
	outFile    string
	verbose    bool
)

// Get top 70 Hacker News stories. If there are no errors, writes to tn_topstories.json.
// Calls the Hacker News API.
//   https://github.com/HackerNews/API
// Uses a Go library to get additional Open Graph data for the article (image, icon, and publisher).
//   https://github.com/otiai10/opengraph
// To Do:

//   - Support using the previous file as a cache for the OG values.
//   - Set timeout on Open Graph fetch.
//   - Setup cron to update data every 10 minutes.
//   - Set user agent when calling URLs.
//   - Add sanitfy check of data.
func main() {
	start := time.Now()

	// Parse command line flags.
	flag.IntVar(&numStories, "num", 70, "number of top stories to get")
	flag.StringVar(&outFile, "out", "hn_topstories.json", "output file JSON")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.Parse()

	if verbose {
		fmt.Printf("Get Hacker News Top Stories\n")
		fmt.Printf("---------------------------\n")
		fmt.Printf("Out File:    %s\n", outFile)
		fmt.Printf("Num Stories: %d\n\n", numStories)
	}

	items, err := hn_og_combo.GetTopStories(numStories)
	if err != nil {
		fmt.Printf("Problem getting top stories: %s\n", err)
		return
	}

	if len(items) < 10 {
		fmt.Printf("Hacker News API returned less than 10 stories, so not writing to %s.\n", outFile)
		return
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Fatalf("Problem marshalling items: %s", err)
		return
	}

	err = os.WriteFile(outFile, data, 0644)
	if err != nil {
		log.Fatalf("Problem saving to file: %s", err)
		return
	}
	if verbose {
		fmt.Println()
		fmt.Printf("Wrote:       %s (%d items, %d bytes).\n", outFile, len(items), len(data))
		fmt.Printf("Took:        %s\n", time.Since(start))
		fmt.Println()
	}
}
