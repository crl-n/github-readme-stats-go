package main

// LanguageStat represents statistics for a programming language in a user's repositories.
// BytesOfCode is the number of bytes of code written in the language.
// Percentage is the percentage of all bytes of code that are written in the language.
type LanguageStat struct {
	BytesOfCode int
	Percentage  float32
}

func combineRepoLanguageStats(repos []Repo) map[string]LanguageStat {
	langStats := make(map[string]LanguageStat)

	// Aggregate bytes of code values, keep track of total
	totalBytes := 0
	for _, repo := range repos {
		for key, bytesOfCode := range repo.Languages {
			langStats[key] = LanguageStat{BytesOfCode: langStats[key].BytesOfCode + bytesOfCode}
			totalBytes += bytesOfCode
		}
	}

	// Count percentages
	for key := range langStats {
		langStat := langStats[key]
		langStat.Percentage = float32(langStat.BytesOfCode) / float32(totalBytes) * 100.0
		langStats[key] = langStat
	}

	return langStats
}
