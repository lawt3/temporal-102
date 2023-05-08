package loanprocess

import (
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/worker"
	"testing"
)

func TestReplayWorkflowHistoryFromFile(t *testing.T) {

	// TODO Create the Workflow Replayer
	replayer := worker.NewWorkflowReplayer()

	// TODO Register the Workflow Definition with the Replayer
	replayer.RegisterWorkflow(LoanProcessingWorkflow)

	// TODO Replay the Event History from the JSON file you downloaded
	err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "history_for_original_execution.json")

	// TODO Use assert.NoError to verify that replaying the history
	//      does not return an error
	assert.NoError(t, err)
}
