package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/okoeth/serverless-demo/commons/pkg/services"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/dghubble/go-twitter/twitter"
)

var target = BuzzwordCounts{
	Keyword: "#cloud",
	Buzzwords: map[string]*BuzzwordCount{
		"#serverless": &BuzzwordCount{
			Keyword:    "#cloud",
			Buzzword:   "#serverless",
			Count:      2,
			LastUpdate: time.Now(),
		},
		"#public": &BuzzwordCount{
			Keyword:    "#cloud",
			Buzzword:   "#serverless",
			Count:      3,
			LastUpdate: time.Now(),
		},
		"#kubernetes": &BuzzwordCount{
			Keyword:    "#cloud",
			Buzzword:   "#serverless",
			Count:      2,
			LastUpdate: time.Now(),
		},
	},
}

var same = BuzzwordCounts{
	Keyword: "#cloud",
	Buzzwords: map[string]*BuzzwordCount{
		"#serverless": &BuzzwordCount{
			Keyword:    "#cloud",
			Buzzword:   "#serverless",
			Count:      2,
			LastUpdate: time.Now(),
		},
		"#cloudfoundry": &BuzzwordCount{
			Keyword:    "#cloud",
			Buzzword:   "#cloudfoundry",
			Count:      3,
			LastUpdate: time.Now(),
		},
		"#kubernetes": &BuzzwordCount{
			Keyword:    "#cloud",
			Buzzword:   "#kubernetes",
			Count:      4,
			LastUpdate: time.Now(),
		},
	},
}

var tweet = twitter.Tweet{
	Text:   "This is #cloud #public #serverless # for all public",
	Source: "#cloud",
}

var anotherTweet = twitter.Tweet{
	Text:   "This is #ai #machinelearning #serverless # for all public",
	Source: "#ai",
}

var empty = BuzzwordCounts{
	Keyword:   "#cloud",
	Buzzwords: make(map[string]*BuzzwordCount),
}

var other = BuzzwordCounts{
	Keyword:   "#other",
	Buzzwords: make(map[string]*BuzzwordCount),
}

var testEvents = []services.Event{
	services.Event{
		ID:        "1",
		Shard:     "1",
		Timestamp: time.Now(),
		Source:    "Poll-Tweet",
		EventType: "Tweet",
		Payload:   &tweet,
	},
	services.Event{
		ID:        "2",
		Shard:     "2",
		Timestamp: time.Now(),
		Source:    "Poll-Tweet",
		EventType: "Tweet",
		Payload:   &anotherTweet,
	},
}

func TestNewBuzzwordCounts(t *testing.T) {
	bc := NewBuzzwordCounts("#cloud")
	if bc.Keyword != "#cloud" {
		t.Errorf("Keyword not properly initialised")
	}
}

func TestAddSameBuzzwordCounts(t *testing.T) {
	AddBuzzwordCounts(&target, &same)
	if len(target.Buzzwords) != 4 {
		t.Errorf("Target structure incomplete")
	}
}

func TestAddOtherBuzzwordCounts(t *testing.T) {
	AddBuzzwordCounts(&target, &other)
	if len(target.Buzzwords) != 3 {
		t.Errorf("Target structure changed")
	}
}

func TestAddToEmptyBuzzwordCounts(t *testing.T) {
	AddBuzzwordCounts(&empty, &same)
	if len(empty.Buzzwords) != 3 {
		t.Errorf("Target structure incomplete")
	}
}

func TestCollectBuzzwords(t *testing.T) {
	b := CollectBuzzwords(testEvents)
	if len(b) != 2 {
		t.Errorf("Incorrect number of buzzwords: got %d, expected 2", len(b))
	}
	if len(b["#cloud"].Buzzwords) != 2 {
		t.Errorf("Incorrect number of buzzword counts: got %d, expected 2", len(b["#cloud"].Buzzwords))
	}
	t.Logf("Result: %+v", b)
}

func TestCollectBuzzwordCounts(t *testing.T) {
	CollectBuzzwordCounts(&tweet, &empty)
	if len(empty.Buzzwords) != 2 {
		t.Errorf("Buzzword count in total incorrect")
	}
	if empty.Buzzwords["#public"].Count != 1 {
		t.Errorf("Buzzword count for public incorrect")
	}
	t.Logf("Result: %+v", empty)
}

func TestProcessEvents(t *testing.T) {
	streamName := os.Getenv("AWS_EVENT_STREAM_NAME")
	if streamName == "" {
		t.Errorf("Stream name not defined, try export AWS_EVENT_STREAM_NAME=...")
		return
	}
	svc := kinesis.New(session.New())
	ds, err := svc.DescribeStream(&kinesis.DescribeStreamInput{
		StreamName: aws.String(streamName),
	})
	if err != nil {
		t.Errorf("Error describing stream: %s", err.Error())
		return
	}
	for _, shard := range ds.StreamDescription.Shards {
		sio, err := svc.GetShardIterator(&kinesis.GetShardIteratorInput{
			ShardId:           shard.ShardId,
			ShardIteratorType: aws.String("TRIM_HORIZON"),
			StreamName:        aws.String(streamName),
		})
		if err != nil {
			t.Errorf("Error creating shard iterator: %s", err.Error())
			return
		}
		si := sio.ShardIterator
		t.Logf("Shard Iterator: %+v", si)
		for i := 0; i < 10; i++ {
			r, err := svc.GetRecords(&kinesis.GetRecordsInput{
				ShardIterator: si,
			})
			if err != nil {
				t.Errorf("Error getting records: %s", err.Error())
				return
			}
			t.Logf("Found number of records: %d", len(r.Records))
			t.Logf("Millies behind: %d", r.MillisBehindLatest)
			if len(r.Records) > 0 {
				for i := range r.Records {
					var event services.Event
					if err := json.Unmarshal(r.Records[i].Data, &event); err != nil {
						t.Errorf("Error unmarshalling event: %s", err.Error())
						return
					}
					t.Logf("Event: %+v\n", event)
					tweet, err := CreateTweetFromMap(event.Payload.(map[string]interface{}))
					if err != nil {
						t.Errorf("Error creating tweet: %s", err.Error())
					}
					t.Logf("Tweet: %+v\n", tweet)
				}
			}
			si = r.NextShardIterator
		}
	}
}
