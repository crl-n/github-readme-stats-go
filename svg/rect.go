package svg

import (
	"encoding/xml"
	"fmt"
)

type RectParams struct {
	Width  int
	Height int
	Fill   string
	Rx     int
	Ry     int
	Stroke string
}

// Represents a rect XML-element
type Rect struct {
	XMLName xml.Name `xml:"rect"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	Fill    string   `xml:"fill,attr"`
	Rx      string   `xml:"rx,attr"`
	Ry      string   `xml:"ry,attr"`
	Stroke  string   `xml:"stroke,attr"`
}

func NewRect(params RectParams) *Rect {
	r := &Rect{
		Width:  fmt.Sprint(params.Width),
		Height: fmt.Sprint(params.Height),
		Fill:   params.Fill,
		Rx:     fmt.Sprint(params.Rx),
		Ry:     fmt.Sprint(params.Ry),
		Stroke: params.Stroke,
	}
	return r
}
