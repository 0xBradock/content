package taskmanager_test

import (
	"context"
	"testing"
	"time"

	"github.com/0xBradock/go-worker/taskmanager"
	"github.com/stretchr/testify/assert"
)

// addParams is required so that the `CreateTask` can infer the task's parameters
type addParams struct {
	// A is the first parameter for the task
	A int

	// B is the second parameter
	B int
}

// addTask is task that will be passed into the worker.
// It requires always 2 parameters:
// - context.Context
// - params: that should be typed
func addTask(ctx context.Context, params addParams) (int, error) {
	time.Sleep(100 * time.Millisecond)
	return params.A + params.B, nil
}

func TestSimpleTaskManager(t *testing.T) {
	repo := taskmanager.NewInMemoryRepo()
	manager := taskmanager.NewTaskManager(repo)
	ctx := context.TODO()

	// The taks creation should be done during application and server configuration.
	// addLauncher should be passed as a dependency to the handler.
	addLauncher := taskmanager.CreateTask(ctx, manager, addTask, taskmanager.TasksOptions{
		Timeout: 1 * time.Second,
	})

	// addLauncher is called in the handler
	taskID, err := addLauncher(addParams{A: 3, B: 5})
	assert.NoError(t, err)
	assert.NotEmpty(t, taskID)

	// Allow some time for goroutine to finish
	time.Sleep(200 * time.Millisecond)

	taskInfo, err := repo.GetTask(taskID)
	assert.NoError(t, err)
	assert.Equal(t, taskmanager.TaskResolved, taskInfo.Status)

	// You could cast Result back to int if your repo/taskmanager stores results
	result, ok := taskInfo.Result.(int)
	assert.True(t, ok)
	assert.Equal(t, 8, result)
}
