package dom

import (
	"bytes"
	"encoding/xml"
	"testing"
)

const sampleXml = `
<?xml version="1.0"?>
<catalog>
   <book id="bk101">
      <author>Gambardella, Matthew</author>
      <title>XML Developer's Guide</title>
      <genre>Computer</genre>
      <price>44.95</price>
      <publish_date>2000-10-01</publish_date>
      <description>An in-depth look at creating applications
      with XML.</description>
   </book>
   <book id="bk102">
      <author>Ralls, Kim</author>
      <title>Midnight Rain</title>
      <genre>Fantasy</genre>
      <price>5.95</price>
      <publish_date>2000-12-16</publish_date>
      <description>A former architect battles corporate zombies,
      an evil sorceress, and her own childhood to become queen
      of the world.</description>
   </book>
   <book id="bk103">
      <author>Corets, Eva</author>
      <title>Maeve Ascendant</title>
      <genre>Fantasy</genre>
      <price>5.95</price>
      <publish_date>2000-11-17</publish_date>
      <description>After the collapse of a nanotechnology
      society in England, the young survivors lay the
      foundation for a new society.</description>
   </book>
</catalog>
`

func TestDocumentFind(t *testing.T) {
	doc, err := buildFromSampleXml()
	if err != nil {
		t.Errorf("error occurred when building DOM: %v", err)
	}

	book := doc.Find(xml.Name{"", "book"})
	// a book element should be found.
	if book == nil {
		t.Errorf("book shouldn't be nil")
	}
	// the first book element should be found.
	if id := book.Attributes[xml.Name{"", "id"}]; id != "bk101" {
		t.Errorf("expected id = %q but got %q", "bk101", id)
	}
}

func TestDocumentList(t *testing.T) {
	doc, err := buildFromSampleXml()
	if err != nil {
		t.Errorf("error occurred when building DOM: %v", err)
	}

	// three book elements should be found.
	books := doc.List(xml.Name{"", "book"})
	if count := len(books); count != 3 {
		t.Errorf("expected len(books) = %v but got %v", 3, count)
	}

	// the book elements should be in order.
	expectedIds := []string{"bk101", "bk102", "bk103"}
	for i, book := range books {
		if id := book.Attributes[xml.Name{"", "id"}]; id != expectedIds[i] {
			t.Errorf("expected id = %q but got %q", expectedIds[i], id)
		}
	}
}

func TestDocumentQueryOne(t *testing.T) {
	doc, err := buildFromSampleXml()
	if err != nil {
		t.Errorf("error occurred when building DOM: %v", err)
	}

	author := doc.QueryOne(xml.Name{"", "book"}, xml.Name{"", "author"})
	if author == nil {
		t.Errorf("author shouldn't be nil")
	}
	if name := author.EnclosedText; name != "Gambardella, Matthew" {
		t.Errorf("expected author name = %q but got %q", "Gambardella, Matthew", name)
	}
}

func TestDocumentQueryMany(t *testing.T) {
	doc, err := buildFromSampleXml()
	if err != nil {
		t.Errorf("error occurred when building DOM: %v", err)
	}

	books := doc.QueryMany(xml.Name{"", "book"})
	if books == nil {
		t.Errorf("books shouldn't be nil")
	}
	if count := len(books); count != 3 {
		t.Errorf("expected len(books) = %v but got %v", 3, count)
	}

	authors := doc.QueryMany(xml.Name{"", "book"}, xml.Name{"", "author"})
	if authors == nil {
		t.Errorf("authors shouldn't be nil")
	}
	if count := len(authors); count != 3 {
		t.Errorf("expected len(authors) = %v but got %v", 3, count)
	}
}

func buildFromSampleXml() (*Document, error) {
	buf := bytes.NewBufferString(sampleXml)
	return BuildFromReader(buf)
}
