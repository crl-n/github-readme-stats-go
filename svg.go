package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	Content string   `xml:",innerxml"`
}

func GenerateTestSVG() {
	svg := &SVG{
		Width:   "100",
		Height:  "100",
		Content: `<circle cx="50" cy="50" r="40" stroke="green" stroke-width="4" fill="yellow" />`,
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
