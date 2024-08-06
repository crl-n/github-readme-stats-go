package svg

import (
	"encoding/xml"
	"strconv"
)

type TextParams struct {
	X       int
	Y       int
	Content string
}

type Text struct {
	XMLName xml.Name `xml:"text"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
	Content string   `xml:",innerxml"`
	Class   string   `xml:"class,attr"`
}

func NewText(params TextParams) *Text {
	return &Text{
		X:       strconv.Itoa(params.X),
		Y:       strconv.Itoa(params.Y),
		Content: params.Content,
	}
}

func (t *Text) SetClass(class string) {
	t.Class = class
}
