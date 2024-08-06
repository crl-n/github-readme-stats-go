package svg

import (
	"fmt"

	"github.com/crl-n/github-readme-stats-go/stats"
)

type LanguageStatsCard struct {
	stats stats.LanguageStats
}

const (
	cardWidth       = 300
	cardHeight      = 285
	cardBgColor     = "#ffffff"
	cardBorderColor = "#e4e2e2"
	paddingX        = 24
	paddingTop      = 24
	langFontSize    = "11px"
	langFontWeight  = 400
	langRowGap      = 150
	font            = "\"Segoe UI\", sans-serif"
	numberOfLangs   = 6
	rowGap          = 35
)

func NewLanguageStatsCard(langStats stats.LanguageStats) LanguageStatsCard {
	return LanguageStatsCard{langStats}
}

func addStyles(svg *SVG) {
	styleElement := Style{}
	styleElement.AppendContent(
		fmt.Sprintf(
			`text { font: %d %s %s }`, langFontWeight, langFontSize, font,
		),
	)
	svg.AppendElement(styleElement)
}

func addCardBackground(svg *SVG) {
	bgRect := NewRect(RectParams{
		Width:  cardWidth,
		Height: cardHeight,
		Fill:   cardBgColor,
		Rx:     4,
		Ry:     4,
		Stroke: cardBorderColor,
	})
	svg.AppendElement(bgRect)
}

func addLanguageRows(svg *SVG, langStats stats.LanguageStats) {
	topLangs := langStats.Top(numberOfLangs)

	for i, stat := range topLangs {
		g := NewGroup()

		y := i*rowGap + paddingTop

		langName := NewText(TextParams{paddingX, y, stat.Language})
		g.AppendElement(langName)

		langStat := NewText(TextParams{
			paddingX + langRowGap,
			y,
			fmt.Sprintf("%.2f %%", stat.Percentage),
		})
		g.AppendElement(langStat)

		svg.AppendElement(g)
	}
}

func (card *LanguageStatsCard) GenerateSVGFile() {
	svg := &SVG{
		Width:    fmt.Sprint(cardWidth),
		Height:   fmt.Sprint(cardHeight),
		Elements: []interface{}{},
	}

	addStyles(svg)
	addCardBackground(svg)
	addLanguageRows(svg, card.stats)

	svg.WriteToFile("langs.svg")
}
