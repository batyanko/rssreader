package reader

import (
	"encoding/xml"
	"net/http"
	"time"
)

// rssBody is root of RSS data structure
type rssBody struct {
	Channel rssChannel `xml:"channel"`
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

// Parse returns a slice of RssItem
func Parse(urls []string) ([]RssItem, error) {
	channels := []rssChannel{}
	items := []RssItem{}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		var body []byte
		resp.Body.Read(body)
		parseRss(body)

		rss, err := parseRss(body)
		if err != nil {
			return nil, err
		}
		channels = append(channels, rss.Channel)
		resp.Body.Close()
	}
	return items, nil
}

// customTime is used to unmarshal RSS pubDate
type customTime struct {
	time.Time
}

// Custom unmarshaler for RSS pubDate format
func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	parse, err := time.Parse(time.RFC1123, v)
	if err != nil {
		// Try another variant of RFC1123
		if parse, err = time.Parse(time.RFC1123Z, v); err != nil {
			return err
		}
	}
	*c = customTime{parse}
	return nil
}

// parseRss into a struct
func parseRss(resp []byte) (rssBody, error) {
	var channel rssBody
	return channel, xml.Unmarshal(resp, &channel)
}
