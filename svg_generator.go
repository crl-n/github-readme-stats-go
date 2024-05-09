package main

import (
	"fmt"
	"strconv"

	"github.com/crl-n/github-readme-stats-go/stats"
)

type SVGGenerator struct{}

func (gen *SVGGenerator) GenerateLangStatsCard(langStats stats.LanguageStats) {
	topLangs := langStats.Top(6)

	svg := &SVG{
		Width:  "300",
		Height: "275",
		Elements: []interface{}{
			Style{Content: `text { font: 400 18px "Segoe UI", sans-serif }`},
			Rect{Width: "300", Height: "275", Fill: "white"},
		},
	}

	for i, stat := range topLangs {
		svg.Elements = append(svg.Elements, Text{X: strconv.Itoa(10), Y: strconv.Itoa(i*35 + 30), Content: stat.Language})
		svg.Elements = append(svg.Elements, Text{X: strconv.Itoa(160), Y: strconv.Itoa(i*35 + 30), Content: fmt.Sprintf("%.2f %%", stat.Percentage)})
	}

	svg.WriteToFile("langs.svg")
}
