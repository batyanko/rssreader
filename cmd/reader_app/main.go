package main

import (
	"fmt"

	"github.com/batyanko/rssreader/pkg/reader"
)

func main() {
	// TODO Add cmd arguments

	args := []string{
		"http://www.rssboard.org/files/sample-rss-2.xml",
		"http://feeds.bbci.co.uk/news/world/europe/rss.xml",
	}

	// TODO group items per channel, sort by pubDate. Display accordingly.

	items, err := reader.Parse(args)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Source: %s\n", item.Source)
		fmt.Printf("Source URL: %s\n", item.SourceURL)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Publish Date: %v\n", item.PublishDate)
		fmt.Printf("Description: %s\n\n", item.Description)
	}
}
