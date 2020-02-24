package reader

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

// Parse returns a slice of RssItem
func Parse(urls []string) ([]RssItem, error) {
	channels, err := extractChannels(urls)
	if err != nil {
		return nil, err
	}

	return extractItems(channels), nil
}

// extractChannels given a slice of URLs
func extractChannels(urls []string) ([]rssChannel, error) {
	channels := []rssChannel{}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		rss, err := parseRss(body)
		if err != nil {
			return nil, err
		}
		channels = append(channels, rss.Channels...)
	}
	return channels, nil
}

// exctractItems from a slice of channels
func extractItems(channels []rssChannel) []RssItem {
	items := []RssItem{}

	for _, channel := range channels {
		for _, item := range channel.Items {
			items = append(items, exportItem(item,
				rssSource{
					Title:     channel.Title,
					SourceURL: channel.Link,
				}))
		}
	}
	return items
}

// Convert internal rssItem to output RssItem.
// Source data is used for items which lack it. This should normally contain channel title and link.
func exportItem(item rssItem, source rssSource) RssItem {
	returnItem := RssItem{
		Title:       item.Title,
		Source:      item.Source.Title,
		SourceURL:   item.Source.SourceURL,
		Link:        item.Link,
		PublishDate: item.PublishDate.Time,
		Description: item.Description,
	}

	// Use custom source data for items that lack it.
	if returnItem.Source == "" {
		returnItem.Source = source.Title
	}
	if returnItem.SourceURL == "" {
		returnItem.SourceURL = source.SourceURL
	}

	return returnItem
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
