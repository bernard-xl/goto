package rss

import (
	"encoding/xml"
	"time"
	"xl/feeds"
	"xl/xml/dom"
)

func item() *feeds.Binding {
	item := feeds.NewBinding(newItem)
	item.Children = map[xml.Name]*feeds.Binding{
		rssName("title"):       feeds.NewBinding(itemTitle),
		rssName("link"):        feeds.NewBinding(itemLink),
		rssName("description"): feeds.NewBinding(itemDescription),
		rssName("author"):      feeds.NewBinding(itemAuthor),
		rssName("category"):    feeds.NewBinding(itemCategory),
		rssName("enclosure"):   feeds.NewBinding(itemEnclosure),
		rssName("guid"):        feeds.NewBinding(itemGuid),
		rssName("pubDate"):     feeds.NewBinding(itemPubDate),
	}
	return item
}

func newItem(d *dom.Document, f *feeds.Feed) error {
	f.Entries = append(f.Entries, &feeds.Entry{})
	return nil
}

func itemTitle(d *dom.Document, f *feeds.Feed) error {
	entry := f.Entries[len(f.Entries)-1]
	entry.Title = d.EnclosedText
	return nil
}

func itemLink(d *dom.Document, f *feeds.Feed) error {
	entry := f.Entries[len(f.Entries)-1]
	entry.Link = d.EnclosedText
	return nil
}

func itemDescription(d *dom.Document, f *feeds.Feed) error {
	entry := f.Entries[len(f.Entries)-1]
	entry.Summary = &feeds.Content{Body: d.EnclosedText}
	return nil
}

func itemAuthor(d *dom.Document, f *feeds.Feed) error {
	entry := f.Entries[len(f.Entries)-1]
	entry.Author = &feeds.Author{Name: d.EnclosedText}
	return nil
}

func itemCategory(d *dom.Document, f *feeds.Feed) error {
	entry := f.Entries[len(f.Entries)-1]
	entry.Categories = append(entry.Categories, d.EnclosedText)
	return nil
}

func itemEnclosure(d *dom.Document, f *feeds.Feed) error {
	entry := f.Entries[len(f.Entries)-1]
	entry.Media = &feeds.Media{Uri: d.Attributes[rssName("url")], Type: d.Attributes[rssName("type")]}
	return nil
}

func itemGuid(d *dom.Document, f *feeds.Feed) error {
	entry := f.Entries[len(f.Entries)-1]
	entry.Id = d.EnclosedText
	return nil
}

func itemPubDate(d *dom.Document, f *feeds.Feed) error {
	entry := f.Entries[len(f.Entries)-1]
	date, err := time.Parse(time.RFC822, d.EnclosedText)
	if err != nil {
		return err
	}
	entry.Date = date
	return nil
}
