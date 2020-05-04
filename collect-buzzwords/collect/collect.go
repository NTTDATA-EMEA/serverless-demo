package main

import (
	"fmt"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"log"
	"strings"
	"time"
)

// BuzzwordCounts holds the buzzwords for a single keyword
type BuzzwordCounts struct {
	Keyword   string                    `json:"keyword"`
	Buzzwords map[string]*BuzzwordCount `json:"buzzwords"`
}

// BuzzwordCount holds the counts for a single keyword/buzzword combo
type BuzzwordCount struct {
	Keyword    string    `json:"keyword"`
	Buzzword   string    `json:"buzzword"`
	Count      int       `json:"count"`
	LastUpdate time.Time `json:"last_update"`
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
		ttx := events[i].Object.TweetText
		bw := strings.ToLower(events[i].Subject.Buzzword)
		if _, ok := b[bw]; !ok {
			b[bw] = NewBuzzwordCounts(bw)
		}
		CollectBuzzwordCounts(ttx, bw, b[bw])
	}
	return b
}

// CollectBuzzwordCounts extracts buzzwords (i.e. hashtags) from tweets and increments counters
func CollectBuzzwordCounts(tw string, bw string, bc *BuzzwordCounts) {
	bwtl := strings.ToLower(bw)
	if bwtl != strings.ToLower(bc.Keyword) {
		return
	}
	words := strings.Fields(tw)
	r := strings.NewReplacer("'s", "", ".", "", ",", "", ";", "", ":", "", "_", "", "-", "",
		"—", "", "`", "", "´", "", "—", "", "\\", "", "\"", "", "!", "", "?", "", "'", "", "…", "")
	for _, word := range words {
		if strings.HasPrefix(word, "#") {
			wordtl := r.Replace(strings.ToLower(word))
			log.Printf("Cleaned: %s -> %s", word, wordtl)
			if len(wordtl) > 1 && wordtl != bwtl {
				if _, ok := bc.Buzzwords[wordtl]; !ok {
					bc.Buzzwords[wordtl] = &BuzzwordCount{
						Keyword:    bc.Keyword,
						Buzzword:   wordtl,
						Count:      0,
						LastUpdate: time.Now(),
					}
				}
				bc.Buzzwords[wordtl].Count++
			}
		}
	}
}

// PublishCollectBuzzwordAggregates publishes events with aggregated values
func PublishCollectBuzzwordAggregates(ep services.EventPublisher, cbs map[string]*BuzzwordCounts) error {
	log.Println("PublishBuzzwordAggregates started...")
	events := make([]CollectEvent, len(cbs))
	i := 0
	for k := range cbs {
		events[i] = CollectEvent{
			EventEnvelope: services.EventEnvelope{
				Event:     services.CollectBuzzwordsAggregated,
				Timestamp: time.Time{},
			},
			Subject: CollectEventSubject{PartitionKey: cbs[k].Keyword, Keyword: cbs[k].Keyword},
			Object:  CollectEventObject{AggregatedBuzzwords: cbs[k]},
		}
		jsn, err := events[i].Marshal()
		if err != nil {
			return fmt.Errorf("publish-collect-buzzword-aggregates.marshal error: %w", err)
		}
		log.Printf("Marshalled Collect Event: %s", jsn)
		i++
	}
	ejsn := make([]services.EventJsoner, len(events))
	for i := range events {
		ejsn[i] = &events[i]
	}
	if err := ep.PublishEvents(ejsn); err != nil {
		return fmt.Errorf("publish-collect-buzzword-aggregates.publish-events error: %w", err)
	}
	log.Println("PublishBuzzwordAggregates finished...")
	return nil
}
