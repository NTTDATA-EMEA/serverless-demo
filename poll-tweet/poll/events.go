package main

import "github.com/okoeth/serverless-demo/commons/pkg/services"

type PollEvent struct {
	services.EventEnvelope
}

func (e PollEvent) GetBuzzword() string {
	return e.Subject.Properties["buzzword"]
}
