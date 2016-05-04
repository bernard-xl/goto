package rss

import (
	"encoding/xml"
	"strconv"
	"time"
	"xl/feeds"
	"xl/xml/dom"
)

func channel() *feeds.Binding {
	channel := feeds.NewBinding(nothing)
	channel.Children = map[xml.Name]*feeds.Binding{
		rssName("title"):          feeds.NewBinding(channelTitle),
		rssName("link"):           feeds.NewBinding(channelLink),
		rssName("description"):    feeds.NewBinding(channelDescription),
		rssName("language"):       feeds.NewBinding(channelLanguage),
		rssName("copyright"):      feeds.NewBinding(channelCopyright),
		rssName("managingEditor"): feeds.NewBinding(channelManagingEditor),
		rssName("pubDate"):        feeds.NewBinding(channelPubDateOrLastBuildDate),
		rssName("lastBuildDate"):  feeds.NewBinding(channelPubDateOrLastBuildDate),
		rssName("category"):       feeds.NewBinding(channelCategory),
		rssName("ttl"):            feeds.NewBinding(channelTtl),
		rssName("image"):          image(),
		rssName("item"):           item(),
	}
	return channel
}

func channelTitle(d *dom.Document, f *feeds.Feed) error {
	f.Title = d.EnclosedText
	return nil
}

func channelLink(d *dom.Document, f *feeds.Feed) error {
	f.Link = d.EnclosedText
	return nil
}

func channelDescription(d *dom.Document, f *feeds.Feed) error {
	f.Summary = &feeds.Content{Body: d.EnclosedText}
	return nil
}

func channelLanguage(d *dom.Document, f *feeds.Feed) error {
	f.Locale = d.EnclosedText
	return nil
}

func channelCopyright(d *dom.Document, f *feeds.Feed) error {
	f.Copyright = d.EnclosedText
	return nil
}

func channelManagingEditor(d *dom.Document, f *feeds.Feed) error {
	f.Author = &feeds.Author{Name: d.EnclosedText}
	return nil
}

func channelPubDateOrLastBuildDate(d *dom.Document, f *feeds.Feed) error {
	date, err := time.Parse(time.RFC822, d.EnclosedText)
	if err != nil {
		return err
	}

	if date.After(f.Date) {
		f.Date = date
	} else {
		f.Date = date
	}
	return nil
}

func channelCategory(d *dom.Document, f *feeds.Feed) error {
	f.Categories = append(f.Categories, d.EnclosedText)
	return nil
}

func channelTtl(d *dom.Document, f *feeds.Feed) error {
	ttl, err := strconv.Atoi(d.EnclosedText)
	if err != nil {
		return err
	}
	f.Ttl = ttl
	return nil
}
