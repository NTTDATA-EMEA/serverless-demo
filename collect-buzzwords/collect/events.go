package main

import "github.com/okoeth/serverless-demo/commons/pkg/services"

type CollectEvent struct {
	services.EventEnvelope
}

type PollEvent struct {
	services.EventEnvelope
}

func (e PollEvent) GetBuzzword() string {
	return e.Subject.Properties["buzzword"]
}

func (e PollEvent) GetTweetText() string {
	return e.Object.Properties["body"].(string)
}
