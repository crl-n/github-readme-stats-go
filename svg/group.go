package svg

import "encoding/xml"

// Represents a g svg-element
type Group struct {
	XMLName  xml.Name      `xml:"g"`
	Elements []interface{} `xml:",innerxml"`
}

func NewGroup() *Group {
	return &Group{
		Elements: []interface{}{},
	}
}

func (g *Group) AppendElement(element interface{}) {
	g.Elements = append(g.Elements, element)
}
