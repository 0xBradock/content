package taskmanager_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/0xBradock/go-worker/taskmanager"
	"github.com/stretchr/testify/assert"
)

type AddParams struct {
	A int
	B int
}

func AddTask(ctx context.Context, params AddParams) (int, error) {
	time.Sleep(100 * time.Millisecond)
	return params.A + params.B, nil
}

// A very simple in-memory repository for testing
type InMemoryRepo struct {
	tasks map[string]*taskmanager.TaskRecord
	mu    sync.RWMutex
}

func TestTaskManager_AddTask(t *testing.T) {
	repo := taskmanager.NewInMemoryRepo()
	manager := taskmanager.NewTaskManager(repo)

	addLauncher := taskmanager.CreateTask(manager, AddTask, taskmanager.TasksOptions{
		Retries: 2,
		Timeout: 2 * time.Second,
	})

	params := AddParams{A: 3, B: 5}

	taskID, err := addLauncher(params)
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
