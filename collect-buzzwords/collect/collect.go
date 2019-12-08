package main

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/okoeth/serverless-demo/commons/pkg/services"

	"github.com/dghubble/go-twitter/twitter"
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
		tweet := events[i].GetTweetText()
		shard := events[i].GetBuzzword()
		if _, ok := b[shard]; !ok {
			b[shard] = NewBuzzwordCounts(shard)
		}
		CollectBuzzwordCounts(tweet, b[tweet.Source])
	}
	return b
}

// CollectBuzzwordCounts extracts buzzwords (i.e. hashtags) from tweets and increments counters
func CollectBuzzwordCounts(tweet *twitter.Tweet, bc *BuzzwordCounts) {
	if tweet.Source != bc.Keyword {
		return
	}
	words := strings.Fields(tweet.Text)
	for _, word := range words {
		if strings.HasPrefix(word, "#") && len(word) > 1 && word != tweet.Source {
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

// CreateTweetFromMap re-marshals a generic map to a proper type
func CreateTweetFromMap(m map[string]interface{}) (*twitter.Tweet, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	var t twitter.Tweet
	if err := json.Unmarshal(b, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
