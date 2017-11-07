package rss

import (
	"io"

	"github.com/devopsfaith/krakend/encoding"
	"github.com/mmcdole/gofeed"
)

func Register() error {
	return encoding.Register(Name, func(_ bool) encoding.Decoder { return NewDecoder() })
}

// Name is the key for the rss encoding
const Name = "rss"

// NewDecoder returns the RSS decoder
func NewDecoder() encoding.Decoder {
	fp := gofeed.NewParser()
	return func(r io.Reader, v *map[string]interface{}) error {
		feed, err := fp.Parse(r)
		if err != nil {
			return err
		}
		*(v) = map[string]interface{}{
			"items":       feed.Items,
			"author":      feed.Author,
			"categories":  feed.Categories,
			"custom":      feed.Custom,
			"copyright":   feed.Copyright,
			"description": feed.Description,
			"type":        feed.FeedType,
			"language":    feed.Language,
			"title":       feed.Title,
			"published":   feed.Published,
			"updated":     feed.Updated,
		}
		if feed.Image != nil {
			(*v)["img_url"] = feed.Image.URL
		}
		return nil
	}
}
