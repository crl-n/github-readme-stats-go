package main

import (
	"sort"

	"github.com/crl-n/github-readme-stats-go/githubclient"
)

// LanguageStat represents statistics for a programming language in a user's repositories.
// BytesOfCode is the number of bytes of code written in the language.
// Percentage is the percentage of all bytes of code that are written in the language.
type LanguageStat struct {
	Language    string
	BytesOfCode int
	Percentage  float32
}

type LanguageStats []LanguageStat

func (stats LanguageStats) Top(n int) []LanguageStat {
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Percentage > stats[j].Percentage
	})
	return stats[:n]
}

func NewLanguageStats(repos []githubclient.Repo) LanguageStats {
	langStats := make(map[string]LanguageStat)

	// Aggregate bytes of code values, keep track of total
	totalBytes := 0
	for _, repo := range repos {
		for key, bytesOfCode := range repo.Languages {
			langStats[key] = LanguageStat{BytesOfCode: langStats[key].BytesOfCode + bytesOfCode}
			totalBytes += bytesOfCode
		}
	}

	// Count percentages and add language names
	for key := range langStats {
		langStat := langStats[key]
		langStat.Language = key
		langStat.Percentage = float32(langStat.BytesOfCode) / float32(totalBytes) * 100.0
		langStats[key] = langStat
	}

	// Turn map of lang stats into slice of lang stats
	langStatsSlice := make([]LanguageStat, 0)
	for key := range langStats {
		langStatsSlice = append(langStatsSlice, langStats[key])
	}

	return langStatsSlice
}
