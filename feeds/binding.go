package feeds

import (
	"encoding/xml"
	"xl/xml/dom"
)

type Binding struct {
	Bind     func(*dom.Document, *Feed) error
	Children map[xml.Name]*Binding
}

func NewBinding(bind func(*dom.Document, *Feed) error) *Binding {
	return &Binding{Bind: bind}
}

func Decode(document *dom.Document, binding *Binding) (*Feed, error) {
	result := &Feed{}
	err := innerDecode(result, document, binding)
	return result, err
}

func innerDecode(f *Feed, d *dom.Document, b *Binding) error {
	err := b.Bind(d, f)
	if err != nil {
		return err
	}

	for _, child := range d.EnclosedElements {
		childBinding, exist := b.Children[child.Name]
		if exist {
			innerDecode(f, child, childBinding)
		}
	}

	return nil
}
