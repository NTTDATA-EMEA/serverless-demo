package main

import (
	"strings"
	"time"
)

// BuzzwordCounts holds the buzzwords for a single keyword
type BuzzwordCounts struct {
	Keyword   string
	Buzzwords map[string]*BuzzwordCount
}

// BuzzwordCount holds the counts for a single keyword/buzzword combo
type BuzzwordCount struct {
	Keyword    string
	Buzzword   string
	Count      int
	LastUpdate time.Time
}

// NewBuzzwordCounts creates new set of buzzword counts
func NewBuzzwordCounts(keyword string) *BuzzwordCounts {
	return &BuzzwordCounts{
		Keyword:   keyword,
		Buzzwords: make(map[string]*BuzzwordCount),
	}
}

// CollectBuzzwords extracts payload from events and calls buzzword collector
func CollectBuzzwords(events []PollEvent) map[string]*BuzzwordCounts {
	b := make(map[string]*BuzzwordCounts)
	for i := range events {
		ttx := events[i].GetTweetText()
		bw := events[i].GetBuzzword()
		if _, ok := b[bw]; !ok {
			b[bw] = NewBuzzwordCounts(bw)
		}
		CollectBuzzwordCounts(ttx, bw, b[bw])
	}
	return b
}

// CollectBuzzwordCounts extracts buzzwords (i.e. hashtags) from tweets and increments counters
func CollectBuzzwordCounts(tw string, bw string, bc *BuzzwordCounts) {
	if bw != bc.Keyword {
		return
	}
	words := strings.Fields(tw)
	for _, word := range words {
		if strings.HasPrefix(word, "#") && len(word) > 1 && word != bw {
			if _, ok := bc.Buzzwords[word]; !ok {
				bc.Buzzwords[word] = &BuzzwordCount{
					Keyword:    bc.Keyword,
					Buzzword:   word,
					Count:      0,
					LastUpdate: time.Now(),
				}
			}
			bc.Buzzwords[word].Count++
		}
	}
}

// AddBuzzwordCounts adds the counts from source to target
func AddBuzzwordCounts(target, source *BuzzwordCounts) {
	if target.Keyword != source.Keyword {
		return
	}
	for k := range source.Buzzwords {
		if _, ok := target.Buzzwords[k]; !ok {
			target.Buzzwords[k] = &BuzzwordCount{
				Keyword:    source.Keyword,
				Buzzword:   k,
				Count:      0,
				LastUpdate: time.Now(),
			}
		}
		target.Buzzwords[k].Count += source.Buzzwords[k].Count
	}
}
