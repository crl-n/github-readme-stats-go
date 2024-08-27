package svg

import (
	"encoding/xml"
	"os"

	"github.com/crl-n/github-readme-stats-go/pkg/logger"
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

func (svg *SVG) AppendElement(element interface{}) {
	svg.Elements = append(svg.Elements, element)
}
