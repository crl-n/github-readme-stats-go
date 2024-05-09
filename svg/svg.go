package svg

import (
	"encoding/xml"
	"os"

	"github.com/crl-n/github-readme-stats-go/logger"
)

type SVG struct {
	XMLName  xml.Name      `xml:"svg"`
	Width    string        `xml:"width,attr"`
	Height   string        `xml:"height,attr"`
	Elements []interface{} `xml:",innerxml"`
}

func (svg SVG) WriteToFile(filename string) {
	output, err := xml.MarshalIndent(svg, "", " ")
	if err != nil {
		logger.Errorf("error: %v\n", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		logger.Errorf("error: %v\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(output))
	if err != nil {
		logger.Errorf("error: %v\n", err)
	}
}

type Group struct {
	XMLName  xml.Name      `xml:"g"`
	Elements []interface{} `xml:",innerxml"`
}

type Rect struct {
	XMLName xml.Name `xml:"rect"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	Fill    string   `xml:"fill,attr"`
	Rx      string   `xml:"rx,attr"`
	Ry      string   `xml:"ry,attr"`
}

type Circle struct {
	XMLName     xml.Name `xml:"circle"`
	Cx          string   `xml:"cx,attr"`
	Cy          string   `xml:"cy,attr"`
	R           string   `xml:"r,attr"`
	Stroke      string   `xml:"stroke,attr"`
	StrokeWidth string   `xml:"stroke-width,attr"`
	Fill        string   `xml:"fill,attr"`
}

type Text struct {
	XMLName xml.Name `xml:"text"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
	Content string   `xml:",innerxml"`
}

type Style struct {
	XMLName xml.Name `xml:"style"`
	Content string   `xml:",innerxml"`
}

func GenerateTestSVG() {
	svg := &SVG{
		Width:  "100",
		Height: "100",
		Elements: []interface{}{
			Rect{
				Width:  "100",
				Height: "100",
				Fill:   "red",
			},
			Circle{
				Cx: "50", Cy: "50", R: "40", Stroke: "green", StrokeWidth: "4", Fill: "yellow",
			},
			Text{
				X:       "10",
				Y:       "50",
				Content: "Hello, world",
			},
		},
	}
	svg.WriteToFile("test.svg")
}
