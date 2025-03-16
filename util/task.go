package util

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	manager = &TaskManager{mu: sync.Mutex{}, registry: make(map[string]*Task)}
)

type TaskManager struct {
	mu       sync.Mutex
	registry map[string]*Task
}

type Task struct {
	Id       string
	stop     func()
	callback func(task *Task)
	mu       sync.Mutex
	interval time.Duration
}

func NewTask(taskId string, cb func(task *Task), interval time.Duration) (*Task, error) {
	_, exists := manager.registry[taskId]
	if exists {
		manager.registry[taskId].Stop()
	}

	manager.mu.Lock()
	defer manager.mu.Unlock()
	task := &Task{Id: taskId, callback: cb, interval: interval}
	manager.registry[taskId] = task
	return task, nil
}

func StopTask(taskId string) error {
	task, exists := manager.registry[taskId]
	if !exists {
		return nil
	}

	task.Stop()

	manager.mu.Lock()
	defer manager.mu.Unlock()
	delete(manager.registry, taskId)
	return nil
}

func (t *Task) Run(ctx context.Context) {
	ctx, t.stop = context.WithCancel(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.loop(ctx):
		}
	}
}

func (t *Task) loop(ctx context.Context) <-chan error {
	errch := make(chan error)
	go func() {
		m := time.NewTicker(t.interval)

		defer func() {
			m.Stop()
			if r := recover(); r != nil {
				errch <- fmt.Errorf("panic with error %v", r)
				close(errch)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				close(errch)
				return
			case <-m.C:
				t.do()
			}
		}
	}()

	return errch
}

func (t *Task) do() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.callback(t) // execute task
}

func (t *Task) Stop() {
	if t.stop != nil {
		t.stop()

		manager.mu.Lock()
		defer manager.mu.Unlock()
		if _, exists := manager.registry[t.Id]; exists {
			delete(manager.registry, t.Id)
		}
	}
}
