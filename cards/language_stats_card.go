package cards

import (
	"fmt"

	. "github.com/crl-n/github-readme-stats-go/stats"
	. "github.com/crl-n/github-readme-stats-go/svg"
)

type LanguageStatsCard struct {
	stats LanguageStats
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
	gapBetweenLangRows      = 30
	gapBetweenTitleAndLangs = 35
	numberOfLangs           = 6
	paddingX                = 24
	paddingTop              = 36
	progressBarHeight       = 8
	progressBarWidth        = 252
	title                   = "Most Used Languages"
	titleFontColor          = "#5FA6EE"
	titleFontSize           = "22px"
	titleFontWeight         = 700
)

func NewLanguageStatsCard(langStats LanguageStats) LanguageStatsCard {
	return LanguageStatsCard{langStats}
}

func addStyles(svg *SVG) {
	styleElement := Style{}
	styleElement.AppendContent(
		fmt.Sprintf(
			`text { font-family: %s }`, font,
		),
	)
	styleElement.AppendContent(
		fmt.Sprintf(
			`.title-text { font-weight: %d; font-size: %s; fill: %s }`,
			titleFontWeight,
			titleFontSize,
			titleFontColor,
		),
	)
	styleElement.AppendContent(
		fmt.Sprintf(
			`.language-text { font-weight: %d; font-size: %s }`,
			langFontWeight,
			langFontSize,
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
	titleText := NewText(TextParams{
		X:       paddingX,
		Y:       paddingTop,
		Content: title,
	})
	titleText.SetClass("title-text")
	svg.AppendElement(titleText)
}

func addLanguageRows(svg *SVG, langStats LanguageStats) {
	topLangs := langStats.Top(numberOfLangs)

	for i, stat := range topLangs {
		g := NewGroup()

		y := i*gapBetweenLangRows + paddingTop + gapBetweenTitleAndLangs

		langName := NewText(TextParams{
			X:       paddingX,
			Y:       y,
			Content: stat.Language,
		})
		langName.SetClass("language-text")
		g.AppendElement(langName)

		langStat := NewText(TextParams{
			X:       paddingX + gapBetweenLangAndStat,
			Y:       y,
			Content: fmt.Sprintf("%.2f %%", stat.Percentage),
		})
		langStat.SetClass("language-text")
		g.AppendElement(langStat)

		svg.AppendElement(g)

		progressBar := NewProgressBar(
			stat.Percentage,
			paddingX,
			y+(gapBetweenLangRows/2)-progressBarHeight,
			progressBarHeight,
			progressBarWidth,
		)
		svg.AppendElement(progressBar.ToElementGroup())
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
