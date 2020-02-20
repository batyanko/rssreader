package reader

import "time"

// internal structs to fit the structure of an RSS response

type rssBody struct {
	Channels []rssChannel `xml:"channel"`
}

type rssChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []rssItem `xml:"item"`
}

type rssItem struct {
	Title       string     `xml:"title"`
	Source      rssSource  `xml:"source"`
	Link        string     `xml:"link"`
	PublishDate customTime `xml:"pubDate"`
	Description string     `xml:"description"`
}

type rssSource struct {
	Title     string `xml:",chardata"`
	SourceURL string `xml:"url,attr"`
}

// RssItem is the return item struct per specification
type RssItem struct {
	Title       string
	Source      string
	SourceURL   string
	Link        string
	PublishDate time.Time
	Description string
}

// customTime is used to unmarshal RSS pubDate
type customTime struct {
	time.Time
}
