package taskmanager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

// TaskFunc defines each task type, were P is its parameter type and R is its respose type.
type TaskFunc[P any, R any] func(ctx context.Context, params P) (R, error)

type TaskLauncher[P any] func(P) (string, error)

// TaskStatus holds the different types of status a task can be
type TaskStatus string

// Task statuses
const (
	TaskPending  TaskStatus = "pending"
	TaskRunning  TaskStatus = "running"
	TaskResolved TaskStatus = "resolved"
	TaskRejected TaskStatus = "rejected"
	TaskCanceled TaskStatus = "canceled"
)

// TaskRecord holds the status for each task in the task manager and repo.
type TaskRecord struct {
	// ID is the task unique id.

	// Status is the status the task is in.
	ID string

	// Status is the status the task is in.
	Status TaskStatus

	// CreatedAt is the time the task was created.
	CreatedAt time.Time

	// UpdatedAt is the time the task was last updated.
	UpdatedAt time.Time

	// Result...
	Result any

	// Error stores the error associated with the task.
	Error string
}

// TaksOptions holds the option for each task
type TasksOptions struct {
	Retries uint
	Timeout time.Duration
}

// TaskRepository declares the methods to be implemented by its repo.
type TaskRepository interface {
	Save(task TaskRecord) error
	Upgrade(task TaskRecord) error
	GetTask(id string) (*TaskRecord, error)
}

// TaskManager holds the configuration and tasks for the main task manager.
type TaskManager struct {
	mu    sync.RWMutex
	tasks map[string]context.CancelFunc
	repo  TaskRepository
}

// NewTaskManager returns the main task manager with its repository.
func NewTaskManager(repo TaskRepository) *TaskManager {
	return &TaskManager{
		repo:  repo,
		tasks: make(map[string]context.CancelFunc),
	}
}

// RunTask is the main executor of a task.
// It calls the task in a goroutine and is able to control retries and cancelation.
func (tm *TaskManager) RunTask(
	ctx context.Context,
	id string,
	run func(ctx context.Context) (any, error),
	opts TasksOptions,
) error {
	ctx, cancel := context.WithCancel(ctx)
	if opts.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, opts.Timeout)
	}

	tm.mu.Lock()
	tm.tasks[id] = cancel
	tm.mu.Unlock()

	now := time.Now()
	task := TaskRecord{
		ID:        id,
		Status:    TaskPending,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := tm.repo.Save(task); err != nil {
		return err
	}

	go func() {
		defer cancel()
		tm.updateStatus(id, TaskRunning)

		var result any
		var err error

		for attempt := 0; attempt <= int(opts.Retries); attempt++ {
			select {
			case <-ctx.Done():
				// TODO: This should be able to cancel the task mid flight.
				// It requires sending a cancel message to the task and getting the answer back.
				// tm.updateStatus(id, TaskCanceled)
				return
				// TODO: This will call the function many times without waiting.
				// There should be a ticker to wait for and it should come from a param.
				// The ticker should implemetn an exp backoff
			default:
				result, err = run(ctx)
				if err == nil {
					tm.updateResult(id, result)
					return
				}
			}
			// TODO: Remove this when the backoff ticker is implemented
			time.Sleep(1 * time.Second)
		}

		tm.updateError(id, fmt.Errorf("task(%s) failed after %d retries: %w", id, opts.Retries, err))
	}()

	return nil
}

// updateResult updates the result of a given task id.
func (tm *TaskManager) updateResult(id string, result any) {
	tm.repo.Upgrade(TaskRecord{
		ID:        id,
		Status:    TaskResolved,
		UpdatedAt: time.Now(),
		Result:    result,
	})
}

// updateError updates the error of a given task id.
func (tm *TaskManager) updateError(id string, err error) {
	tm.repo.Upgrade(TaskRecord{
		ID:        id,
		Status:    TaskRejected,
		UpdatedAt: time.Now(),
		Error:     err.Error(),
	})
}

// updateStatus updates the status of a given task id.
func (tm *TaskManager) updateStatus(id string, status TaskStatus) {
	tm.repo.Upgrade(TaskRecord{
		ID:        id,
		Status:    status,
		UpdatedAt: time.Now(),
	})
}

// CreateTask wraps a given task with SubmitTask.
func CreateTask[P any, R any](
	ctx context.Context,
	tm *TaskManager,
	fn TaskFunc[P, R],
	opts TasksOptions,
) func(params P) (string, error) {
	return func(params P) (string, error) {
		uuid, err := uuid.NewV7()
		if err != nil {
			return "", err
		}
		id := uuid.String()

		run := func(ctx context.Context) (any, error) {
			return fn(ctx, params)
		}

		err = tm.RunTask(ctx, id, run, opts)
		if err != nil {
			return "", err
		}

		return id, nil
	}
}
