package svg

import (
	"fmt"
	"strconv"

	"github.com/crl-n/github-readme-stats-go/stats"
)

type SVGGenerator struct{}

const (
	cardWidth    = 300
	cardHeight   = 285
	paddingX     = 24
	paddingTop   = 24
	langFontSize = "11px"
	// titleFontSize = "18px"
	numberOfLangs = 6
	rowGap        = 35
)

func addStyles(svg *SVG) {
	styleElement := Style{Content: fmt.Sprintf(`text { font: 400 %s "Segoe UI", sans-serif }`, langFontSize)}
	svg.Elements = append(svg.Elements, styleElement)
}

func addCardBackground(svg *SVG) {
	bgRect := Rect{Width: fmt.Sprint(cardWidth), Height: fmt.Sprint(cardHeight), Fill: "white", Rx: "4", Ry: "4", Stroke: "#e4e2e2"}
	svg.Elements = append(svg.Elements, bgRect)
}

func addLanguageRows(svg *SVG, langStats stats.LanguageStats) {
	topLangs := langStats.Top(numberOfLangs)

	for i, stat := range topLangs {
		g := &Group{
			Elements: []interface{}{},
		}

		y := i*rowGap + paddingTop

		g.Elements = append(g.Elements, Text{X: strconv.Itoa(paddingX), Y: strconv.Itoa(y), Content: stat.Language})
		g.Elements = append(g.Elements, Text{X: strconv.Itoa(paddingX + 150), Y: strconv.Itoa(y), Content: fmt.Sprintf("%.2f %%", stat.Percentage)})
		svg.Elements = append(svg.Elements, g)
	}
}

func (gen *SVGGenerator) GenerateLangStatsCard(langStats stats.LanguageStats) {
	svg := &SVG{
		Width:    fmt.Sprint(cardWidth),
		Height:   fmt.Sprint(cardHeight),
		Elements: []interface{}{},
	}

	addStyles(svg)
	addCardBackground(svg)
	addLanguageRows(svg, langStats)

	svg.WriteToFile("langs.svg")
}
