package main

import (
	"fmt"
	"sort"
	"strconv"
)

type SVGGenerator struct {
}

func (gen *SVGGenerator) GenerateLangStatsCard(stats []LanguageStat) {
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Percentage > stats[j].Percentage
	})

	svg := &SVG{
		Width:  "300",
		Height: "275",
		Elements: []interface{}{
			Style{Content: `text { font: 400 18px "Segoe UI", sans-serif }`},
			Rect{Width: "300", Height: "275", Fill: "white"},
		},
	}

	for i, stat := range stats {
		if i > 5 {
			break
		}
		fmt.Println(stat)
		svg.Elements = append(svg.Elements, Text{X: strconv.Itoa(10), Y: strconv.Itoa(i*35 + 30), Content: stat.Language})
		svg.Elements = append(svg.Elements, Text{X: strconv.Itoa(160), Y: strconv.Itoa(i*35 + 30), Content: fmt.Sprintf("%.2f %%", stat.Percentage)})
	}

	svg.WriteToFile("langs.svg")
}
