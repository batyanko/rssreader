package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/batyanko/rssreader/pkg/reader"
)

func main() {
	appName := filepath.Base(os.Args[0])
	urls := os.Args[1:]
	help := flag.Bool("h", false, "show this help screen")
	flag.Parse()

	if *help || len(urls) == 0 {
		fmt.Printf("Usage of %s:\n", appName)
		fmt.Printf("Use this tool to parse and display RSS content.\n\n")
		fmt.Printf("Arguments:\n")
		fmt.Printf("Only URL addresses are accepted as arguments.\n\n")
		fmt.Printf("Output order:\n")
		fmt.Printf("Results are displayed per input order of the URLs, then arranged by RSS channel.\n")
		fmt.Printf("RSS items for each channel are sorted by publishing date in descending order.\n")
		return
	}

	for i, url := range urls {
		fmt.Println(i, url)

		items, err := reader.Parse(urls)
		if err != nil {
			// TODO add logger?
			panic(err)
		}

		channels, sortedItems := sortItems(items)

		for _, channel := range channels {
			printItems(channel, sortedItems[channel])
		}
	}
}

func sortItems(items []reader.RssItem) ([]string, map[string][]reader.RssItem) {
	channels := []string{}
	channelItems := map[string][]reader.RssItem{}

	// Populate a map of items for eachChannel
	for _, item := range items {
		// Init if not existing
		if _, ok := channelItems[item.Source]; !ok {
			channels = append(channels, item.Source)
			channelItems[item.Source] = []reader.RssItem{item}
		} else {
			// Append otherwise
			channelItems[item.Source] = append(channelItems[item.Source], item)
		}
	}

	// Sort channels alphabetically
	sort.Strings(channels)

	// Sort items for each channel
	for _, channelItems := range channelItems {
		sort.Slice(channelItems, func(i, j int) bool {
			return channelItems[i].PublishDate.After(channelItems[j].PublishDate)
		})
	}
	return channels, channelItems
}

// Print out a slice of RSS items for a given channel.
func printItems(channel string, items []reader.RssItem) {
	fmt.Printf("-----------------------\n")
	fmt.Printf("RSS items in channel %s:\n", channel)
	fmt.Printf("-----------------------\n\n")
	for _, item := range items {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Source: %s\n", item.Source)
		fmt.Printf("Source URL: %s\n", item.SourceURL)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Publish Date: %v\n", item.PublishDate)
		fmt.Printf("Description: %s\n\n", item.Description)
	}
}