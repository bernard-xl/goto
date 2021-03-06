package rss

import (
	"encoding/xml"

	"github.com/bernard-xl/goto/feeds"
	"github.com/bernard-xl/goto/xml/dom"
)

var Binding *feeds.Binding = rss()

func rss() *feeds.Binding {
	rss := feeds.NewBinding(nothing)
	rss.Children = map[xml.Name]*feeds.Binding{
		rssName("channel"): channel(),
	}
	return rss
}

func rssName(name string) xml.Name {
	return xml.Name{"", name}
}

func nothing(d *dom.Document, f *feeds.Feed) error {
	return nil
}
