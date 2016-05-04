package feeds

import "time"

type Feed struct {
	Locale     string
	Title      string
	Subtitle   string
	Link       string
	Date       time.Time
	Image      string
	Copyright  string
	Author     *Author
	Summary    *Content
	Categories []string
	Entries    []*Entry
	Ttl        int
}

type Entry struct {
	Id         string
	Title      string
	Subtitle   string
	Link       string
	Date       time.Time
	Image      string
	Media      *Media
	Author     *Author
	Summary    *Content
	Categories []string
}

type Author struct {
	Name  string
	Uri   string
	Email string
}

type Content struct {
	Body string
	Type string
}

type Media struct {
	Uri  string
	Type string
}
