package services

import "time"

// Document is a wrapper for a generic data structure
type Document struct {
	ID           string
	Shard        string
	Timestamp    time.Time
	DocumentType string
	DocumentData interface{}
}

// DocumentDB is a cloud provider independent interface to a document
// oriented database like DynamoDB
type DocumentDB interface {
	PutDocument(doc *Document) error
	GetDocument(key string) (*Document, error)
	DeleteDocument(key string) error
}
