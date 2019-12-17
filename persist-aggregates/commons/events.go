package commons

import (
	"encoding/json"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"time"
)

type CollectEvent struct {
	services.EventEnvelope
	Subject CollectEventSubject `json:"subject"`
	Object  CollectEventObject  `json:"object"`
}

type CollectEventSubject struct {
	PartitionKey string
	Keyword      string
}

type CollectEventObject struct {
	AggregatedBuzzwords BuzzwordCounts `json:"aggregated_buzzwords"`
}

func (ce CollectEvent) GetPartitionKey() string {
	return ce.Subject.PartitionKey
}

func (ce CollectEvent) Marshal() ([]byte, error) {
	return json.Marshal(ce)
}

func (ce *CollectEvent) Unmarshal(jsn []byte) error {
	return json.Unmarshal(jsn, ce)
}

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
