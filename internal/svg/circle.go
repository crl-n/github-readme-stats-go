package svg

import "encoding/xml"

type Circle struct {
	XMLName     xml.Name `xml:"circle"`
	Cx          string   `xml:"cx,attr"`
	Cy          string   `xml:"cy,attr"`
	R           string   `xml:"r,attr"`
	Stroke      string   `xml:"stroke,attr"`
	StrokeWidth string   `xml:"stroke-width,attr"`
	Fill        string   `xml:"fill,attr"`
}
