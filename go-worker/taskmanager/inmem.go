package taskmanager

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// In-memory repository for simplicity
type InMemoryRepo struct {
	mu    sync.Mutex
	tasks map[string]TaskRecord
}

// GetTask implements TaskRepository.
func (r *InMemoryRepo) GetTask(id string) (*TaskRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rec, ok := r.tasks[id]
	if !ok {
		return nil, fmt.Errorf("failed to retrieve task with id %s", id)
	}

	return &rec, nil
}

func NewInMemoryRepo() TaskRepository {
	return &InMemoryRepo{
		tasks: make(map[string]TaskRecord),
	}
}

func (r *InMemoryRepo) Save(task TaskRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryRepo) Upgrade(task TaskRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	existing, ok := r.tasks[task.ID]
	if !ok {
		return errors.New("task not found")
	}
	existing.Status = task.Status
	existing.Result = task.Result
	existing.Error = task.Error
	existing.UpdatedAt = time.Now()
	r.tasks[task.ID] = existing
	return nil
}
