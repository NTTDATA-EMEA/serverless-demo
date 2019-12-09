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

	"github.com/stretchr/testify/suite"
)

type CollectBuzzwordsTestSuite struct {
	suite.Suite
	target               BuzzwordCounts
	same                 BuzzwordCounts
	tweet                twitter.Tweet
	tweetText            string
	tweetBuzzword        string
	anotherTweetText     string
	anotherTweetBuzzword string
	empty                BuzzwordCounts
	other                BuzzwordCounts
	testEvents           []PollEvent
}

func (suite *CollectBuzzwordsTestSuite) SetupTest() {
	suite.target = BuzzwordCounts{
		Keyword: "#cloud",
		Buzzwords: map[string]*BuzzwordCount{
			"#serverless": {
				Keyword:    "#cloud",
				Buzzword:   "#serverless",
				Count:      2,
				LastUpdate: time.Now(),
			},
			"#cloudfoundry": {
				Keyword:    "#cloud",
				Buzzword:   "#cloudfoundry",
				Count:      3,
				LastUpdate: time.Now(),
			},
			"#kubernetes": {
				Keyword:    "#cloud",
				Buzzword:   "#kubernetes",
				Count:      2,
				LastUpdate: time.Now(),
			},
		},
	}

	suite.same = BuzzwordCounts{
		Keyword: "#cloud",
		Buzzwords: map[string]*BuzzwordCount{
			"#serverless": {
				Keyword:    "#cloud",
				Buzzword:   "#serverless",
				Count:      2,
				LastUpdate: time.Now(),
			},
			"#cloudfoundry": {
				Keyword:    "#cloud",
				Buzzword:   "#cloudfoundry",
				Count:      3,
				LastUpdate: time.Now(),
			},
			"#kubernetes": {
				Keyword:    "#cloud",
				Buzzword:   "#kubernetes",
				Count:      4,
				LastUpdate: time.Now(),
			},
		},
	}

	suite.tweetText = "This is #cloud #public #serverless # for all public"
	suite.tweetBuzzword = "#cloud"

	suite.anotherTweetText = "This is #ai #machinelearning #serverless # for all public"
	suite.anotherTweetBuzzword = "#ai"

	suite.empty = BuzzwordCounts{
		Keyword:   "#cloud",
		Buzzwords: make(map[string]*BuzzwordCount),
	}

	suite.other = BuzzwordCounts{
		Keyword:   "#other",
		Buzzwords: make(map[string]*BuzzwordCount),
	}

	suite.testEvents = []PollEvent{
		{
			services.EventEnvelope{
				Event:     "",
				Timestamp: time.Now(),
				Object: services.EventObject{
					Id:         "1",
					Name:       "query",
					Properties: map[string]string{},
				},
				Subject: services.EventSubject{
					Id:         "1",
					Name:       "tweet",
					Properties: map[string]string{},
				},
			},
		},
		{
			services.EventEnvelope{
				Event:     "",
				Timestamp: time.Now(),
				Object: services.EventObject{
					Id:         "2",
					Name:       "query",
					Properties: map[string]string{},
				},
				Subject: services.EventSubject{
					Id:         "2",
					Name:       "tweet",
					Properties: map[string]string{},
				},
			},
		},
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestCollectBuzzwordsTestSuite(t *testing.T) {
	suite.Run(t, new(CollectBuzzwordsTestSuite))
}

func (suite *CollectBuzzwordsTestSuite) TestNewBuzzwordCounts() {
	t := suite.T()
	bc := NewBuzzwordCounts("#cloud")
	if bc.Keyword != "#cloud" {
		t.Errorf("Keyword not properly initialised")
	}
}

func (suite *CollectBuzzwordsTestSuite) TestAddSameBuzzwordCounts() {
	t := suite.T()
	AddBuzzwordCounts(&suite.target, &suite.same)
	if len(suite.target.Buzzwords) != 3 {
		t.Errorf("Target structure incomplete")
	}
}

func (suite *CollectBuzzwordsTestSuite) TestAddOtherBuzzwordCounts() {
	t := suite.T()
	AddBuzzwordCounts(&suite.target, &suite.other)
	if len(suite.target.Buzzwords) != 3 {
		t.Errorf("Target structure changed")
	}
}

func (suite *CollectBuzzwordsTestSuite) TestAddToEmptyBuzzwordCounts() {
	t := suite.T()
	AddBuzzwordCounts(&suite.empty, &suite.same)
	if len(suite.empty.Buzzwords) != 3 {
		t.Errorf("Target structure incomplete")
	}
}

func (suite *CollectBuzzwordsTestSuite) TestCollectBuzzwords() {
	t := suite.T()
	b := CollectBuzzwords(suite.testEvents)
	if len(b) != 2 {
		t.Errorf("Incorrect number of buzzwords: got %d, expected 2", len(b))
	}
	if len(b["#cloud"].Buzzwords) != 2 {
		t.Errorf("Incorrect number of buzzword counts: got %d, expected 2", len(b["#cloud"].Buzzwords))
	}
	t.Logf("Result: %+v", b)
}

func (suite *CollectBuzzwordsTestSuite) TestCollectBuzzwordCounts() {
	t := suite.T()
	CollectBuzzwordCounts(suite.tweetText, suite.tweetBuzzword, &suite.empty)
	if len(suite.empty.Buzzwords) != 2 {
		t.Errorf("Buzzword count in total incorrect")
	}
	if suite.empty.Buzzwords["#public"].Count != 1 {
		t.Errorf("Buzzword count for public incorrect")
	}
	t.Logf("Result: %+v", suite.empty)
}

func (suite *CollectBuzzwordsTestSuite) TestProcessEvents() {
	t := suite.T()
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
					var event PollEvent
					if err := json.Unmarshal(r.Records[i].Data, &event); err != nil {
						t.Errorf("Error unmarshalling event: %s", err.Error())
						return
					}
					t.Logf("Event: %+v\n", event.GetTweetText())
				}
			}
			si = r.NextShardIterator
		}
	}
}
