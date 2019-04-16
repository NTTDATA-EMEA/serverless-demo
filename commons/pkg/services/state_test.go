package services

import (
	"os"
	"testing"
)

func getStateStorers() []StateStorer {
	var stateStorer []StateStorer
	stateStorer =
		append(stateStorer, NewAwsStateStorer(os.Getenv("TWITTER_STATE_BUCKET"), "testState.json"))
	stateStorer =
		append(stateStorer, NewLocalStateStorer("/tmp", "testState.json"))
	return stateStorer
}

func getTestState() State {
	testState := make(State)
	testState["#cloud"] = 0
	testState["#ai"] = 0
	testState["#iot"] = 0
	return testState
}

func TestStateStorer(t *testing.T) {
	for _, stateStorer := range getStateStorers() {
		if err := stateStorer.SetState(getTestState()); err != nil {
			t.Errorf("Error saving state: %s", err.Error())
		}
		state, err := stateStorer.GetState()
		if err != nil {
			t.Errorf("Error loading state: %s", err.Error())
		}
		t.Logf("Retrieved state: %+v", state)
		if err := stateStorer.DeleteState(); err != nil {
			t.Errorf("Error deleting state: %s", err.Error())
		}
	}
}
