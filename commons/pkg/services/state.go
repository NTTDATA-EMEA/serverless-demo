package services

// State for this service is comprised of a set of Twitter query strings
// and the latest ID found (Since_ID).
type State map[string]int64

// StateStorer is a cloud provider independent interface to store state
// in serverless services like AWS S3
type StateStorer interface {
	GetState() (State, error)
	SetState(state State) error
	DeleteState() error
}
