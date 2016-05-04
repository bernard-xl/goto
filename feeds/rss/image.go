package rss

import (
	"encoding/xml"

	"github.com/bernard-xl/goto/feeds"
	"github.com/bernard-xl/goto/xml/dom"
)

func image() *feeds.Binding {
	image := feeds.NewBinding(nothing)
	image.Children = map[xml.Name]*feeds.Binding{
		rssName("url"): feeds.NewBinding(imageUrl),
	}
	return image
}

func imageUrl(d *dom.Document, f *feeds.Feed) error {
	f.Image = d.EnclosedText
	return nil
}
