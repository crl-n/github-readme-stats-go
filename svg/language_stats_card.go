package svg

import (
	"fmt"

	"github.com/crl-n/github-readme-stats-go/stats"
)

type LanguageStatsCard struct {
	stats stats.LanguageStats
}

const (
	cardWidth               = 300
	cardHeight              = 285
	cardBgColor             = "#ffffff"
	cardBorderColor         = "#e4e2e2"
	font                    = "\"Segoe UI\", sans-serif"
	langFontSize            = "11px"
	langFontWeight          = 400
	gapBetweenLangAndStat   = 150
	gapBetweenLangRows      = 35
	gapBetweenTitleAndLangs = 50
	numberOfLangs           = 6
	paddingX                = 24
	paddingTop              = 24
	title                   = "Most Used Languages"
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

func addTitle(svg *SVG) {
	titleText := NewText(TextParams{paddingX, paddingTop, title})
	svg.AppendElement(titleText)
}

func addLanguageRows(svg *SVG, langStats stats.LanguageStats) {
	topLangs := langStats.Top(numberOfLangs)

	for i, stat := range topLangs {
		g := NewGroup()

		y := i*gapBetweenLangRows + paddingTop + gapBetweenTitleAndLangs

		langName := NewText(TextParams{paddingX, y, stat.Language})
		g.AppendElement(langName)

		langStat := NewText(TextParams{
			paddingX + gapBetweenLangAndStat,
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
	addTitle(svg)
	addLanguageRows(svg, card.stats)

	svg.WriteToFile("langs.svg")
}
