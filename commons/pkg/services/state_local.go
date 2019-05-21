package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// LocalStateStorer implements the StateStorer interface
type LocalStateStorer struct {
	Address  string
	Filename string
}

// NewLocalStateStorer returns a StateStorer implementation
func NewLocalStateStorer(address, filename string) StateStorer {
	return &LocalStateStorer{
		Address:  address,
		Filename: filename,
	}
}

// GetState retrieves the state from the S3 bucket
func (as *LocalStateStorer) GetState() (State, error) {
	if as.Address == "" || as.Filename == "" {
		return nil, errors.New("storer address and/or filename required")
	}
	buffer, err := ioutil.ReadFile(as.Address + "/" + as.Filename)
	if err != nil {
		return nil, err
	}
	state := make(map[string]int64)
	err = json.Unmarshal(buffer, &state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

// SetState writes the state to the S3 bucket
func (as *LocalStateStorer) SetState(state State) error {
	json, err := json.Marshal(state)
	if err != nil {
		return err
	}
	fmt.Printf("Serialised testState: %s\n", json)
	if err := ioutil.WriteFile(as.Address+"/"+as.Filename, json, 0644); err != nil {
		return err
	}
	return nil
}

// DeleteState writes the state to the S3 bucket
func (as *LocalStateStorer) DeleteState() error {
	if err := os.Remove(as.Address + "/" + as.Filename); err != nil {
		return err
	}
	return nil
}
