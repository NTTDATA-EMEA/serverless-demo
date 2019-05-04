package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LocalDocumentDB implements the DocumentDB interface
type LocalDocumentDB struct {
	Path string
}

// NewLocalDocumentDB returns a DocumentDB implementation
func NewLocalDocumentDB(path string) DocumentDB {
	return &LocalDocumentDB{
		Path: path,
	}
}

// GetDocument retrieves the document from the local filesystem
func (db *LocalDocumentDB) GetDocument(key string) (*Document, error) {
	buffer, err := ioutil.ReadFile(db.Path + "/" + key + ".db")
	if err != nil {
		return nil, err
	}
	var doc Document
	err = json.Unmarshal(buffer, &doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

// PutDocument writes the document to the local filesystem
func (db *LocalDocumentDB) PutDocument(doc *Document) error {
	json, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	fmt.Printf("Serialised testState: %s\n", json)
	if err := ioutil.WriteFile(db.Path+"/"+doc.ID+".db", json, 0644); err != nil {
		return err
	}
	return nil
}

// DeleteDocument removes the document from the local filesystem
func (db *LocalDocumentDB) DeleteDocument(key string) error {
	if err := os.Remove(db.Path + "/" + key + ".db"); err != nil {
		return err
	}
	return nil
}
