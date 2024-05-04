package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type SVG struct {
	XMLName  xml.Name      `xml:"svg"`
	Width    string        `xml:"width,attr"`
	Height   string        `xml:"height,attr"`
	Elements []interface{} `xml:",innerxml"`
}

type Rect struct {
	XMLName xml.Name `xml:"rect"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	Fill    string   `xml:"fill,attr"`
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
		},
	}

	output, err := xml.MarshalIndent(svg, "", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	file, err := os.Create("test.svg")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(output))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
