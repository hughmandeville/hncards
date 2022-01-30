package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hughmandeville/hnui/pkg/github"
	"github.com/hughmandeville/hnui/pkg/hn_og_combo"
)

const (
	filePath = "client/public/hn_topstories.json"
)

var (
	numStories int
	out        string
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
	flag.StringVar(&out, "out", "github", "output location (file or github)")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.Parse()

	if out == "github" && os.Getenv("GITHUB_TOKEN") == "" {
		fmt.Printf("GITHUB_TOKEN environment variable must when outputing to GitHub.\n")
	}

	if verbose {
		fmt.Printf("Get Hacker News Top Stories\n")
		fmt.Printf("---------------------------\n")
		fmt.Printf("Output To:   %s\n", out)
		fmt.Printf("Num Stories: %d\n\n", numStories)
	}

	items, err := hn_og_combo.GetTopStories(numStories)
	if err != nil {
		fmt.Printf("Problem getting top stories: %s\n", err)
		return
	}

	if len(items) < 10 {
		fmt.Printf("Hacker News API returned less than 10 stories, so not writing to %s.\n", filePath)
		return
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Fatalf("Problem marshalling items: %s", err)
		return
	}
	fmt.Println()
	if out == "github" {
		ghc := github.NewGitHubController(os.Getenv("GITHUB_TOKEN"), "main", "hughmandeville", "hnui")
		_, sha, err := ghc.GetFile(filePath)
		if err != nil {
			fmt.Printf("Problem getting data from GitHub: %s\n", err)
		} else {
			err = ghc.PutFile(filePath, data, sha)
			if err != nil {
				fmt.Printf("Problem committing data to GitHub: %s\n", err)
			} else {
				fmt.Printf("Committed:   %s (%d items, %d bytes).\n", filePath, len(items), len(data))
			}
		}

	} else {
		err = os.WriteFile(filePath, data, 0644)
		if err != nil {
			log.Fatalf("Problem saving to file: %s", err)
			return
		}
		fmt.Printf("Wrote:       %s (%d items, %d bytes).\n", filePath, len(items), len(data))
	}
	if verbose {

		fmt.Printf("Took:        %s\n", time.Since(start))
		fmt.Println()
	}
}
