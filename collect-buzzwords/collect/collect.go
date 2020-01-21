package main

import (
	"fmt"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	log "github.com/sirupsen/logrus"
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
		bw := events[i].Subject.Buzzword
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
		if strings.HasPrefix(word, "#") && len(word) > 1 && strings.ToLower(word) != strings.ToLower(bw) {
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

// PublishCollectBuzzwordAggregates publishes events with aggregated values
func PublishCollectBuzzwordAggregates(ep services.EventPublisher, cbs map[string]*BuzzwordCounts) error {
	log.Info("PublishBuzzwordAggregates started...")
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
		log.WithField("buzzword", fmt.Sprintf("%s", jsn)).Info("Marshalled Collect Event...")
		i++
	}
	ejsn := make([]services.EventJsoner, len(events))
	for i := range events {
		ejsn[i] = &events[i]
	}
	if err := ep.PublishEvents(ejsn); err != nil {
		return fmt.Errorf("publish-collect-buzzword-aggregates.publish-events error: %w", err)
	}
	log.Info("PublishBuzzwordAggregates finished...")
	return nil
}
