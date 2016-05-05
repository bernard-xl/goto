package dom

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Document struct {
	Name             xml.Name
	Attributes       map[xml.Name]string
	Children         []interface{}
	EnclosedElements []*Document
	EnclosedText     string
}

func BuildFromReader(reader io.Reader) (*Document, error) {
	decoder := xml.NewDecoder(reader)
	stack := make([]*Document, 0, 8)

	for {
		untypedToken, err := decoder.Token()
		if untypedToken == nil && err == io.EOF {
			return stack[0], nil
		}
		if err != nil {
			return nil, err
		}

		switch token := untypedToken.(type) {
		case xml.StartElement:
			doc := new(Document)
			doc.Name = token.Name
			doc.Attributes = decodeAttributes(token.Attr)
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, doc)
				parent.EnclosedElements = append(parent.EnclosedElements, doc)
			}
			stack = append(stack, doc)
		case xml.CharData:
			if len(stack) > 0 {
				current := stack[len(stack)-1]
				text := strings.TrimSpace(string(token))
				if text != "" {
					current.Children = append(current.Children, text)
					current.EnclosedText = current.EnclosedText + text
				}
			}
		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		}
	}
}

func (self *Document) Find(name xml.Name) *Document {
	for _, untypedChild := range self.Children {
		switch child := untypedChild.(type) {
		case *Document:
			if child.Name == name {
				return child
			}
		}
	}
	return nil
}

func (self *Document) List(name xml.Name) []*Document {
	result := make([]*Document, 0, 8)
	for _, untypedChild := range self.Children {
		switch child := untypedChild.(type) {
		case *Document:
			if child.Name == name {
				result = append(result, child)
			}
		}
	}
	return result
}

func (self *Document) QueryOne(path ...xml.Name) *Document {
	result := self

	for len(path) != 0 {
		result = result.Find(path[0])
		path = path[1:]
		if result == nil {
			return nil
		}
	}
	return result
}

func (self *Document) QueryMany(path ...xml.Name) []*Document {
	if len(path) == 0 {
		return nil
	}

	var result []*Document
	target := path[0]
	nextPath := path[1:]

	for _, child := range self.EnclosedElements {
		if child.Name == target {
			if len(path) == 1 {
				result = append(result, self)
			} else {
				subresult := child.QueryMany(nextPath...)
				result = append(result, subresult...)
			}
		}
	}
	return result
}

func (self *Document) String() string {
	buffer := new(bytes.Buffer)
	makeString(buffer, self, 0)
	return buffer.String()
}

func makeString(buffer *bytes.Buffer, doc *Document, depth int) {
	fmt.Fprintf(buffer, "%*v%v ", depth*2, "", doc.Name.Local)

	if len(doc.Attributes) == 0 {
		fmt.Fprint(buffer, "\n")
	} else {
		fmt.Fprint(buffer, "=>")
		for key, value := range doc.Attributes {
			fmt.Fprintf(buffer, " %v: %v", key, value)
		}
		fmt.Fprint(buffer, "\n")
	}

	for _, untypedChild := range doc.Children {
		switch child := untypedChild.(type) {
		case *Document:
			makeString(buffer, child, depth+1)
		case string:
			if len(child) > 77 {
				fmt.Fprintf(buffer, "%*v%v...\n", (depth+1)*2, "", child[:77])
			} else {
				fmt.Fprintf(buffer, "%*v%v\n", (depth+1)*2, "", child)
			}
		}
	}
}

func decodeAttributes(attributes []xml.Attr) map[xml.Name]string {
	result := make(map[xml.Name]string)
	for _, attr := range attributes {
		result[attr.Name] = attr.Value
	}
	return result
}
