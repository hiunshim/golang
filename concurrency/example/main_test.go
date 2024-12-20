package main

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNoChannelExample(t *testing.T) {
	go func() {
		fmt.Println("Running without channels")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("No-channel example done")
}

func TestChannelExample(t *testing.T) {
	ch := make(chan string)
	go func() { ch <- "Running with goroutine" }()
	fmt.Println(<-ch)
	fmt.Println("Channel example done")
}

func TestWorkerPool_Success(t *testing.T) {
	tasks := []Task{
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
	}

	workerCount := 2
	errors := WorkerPool(tasks, workerCount)

	if len(errors) != 0 {
		t.Fatalf("expected no errors, got %d errors: %v", len(errors), errors)
	}
}

func TestWorkerPool_Errors(t *testing.T) {
	// Create tasks with some failures
	tasks := []Task{
		func() error { return errors.New("Task 1 failed") },
		func() error { return nil },
		func() error { return errors.New("Task 3 failed") },
	}

	workerCount := 3
	errors := WorkerPool(tasks, workerCount)

	// Verify the correct number of errors
	if len(errors) != 2 {
		t.Fatalf("expected 2 errors, got %d: %v", len(errors), errors)
	}
}

func TestWorkerPool_Concurrency(t *testing.T) {
	var mu sync.Mutex
	taskOrder := []int{}

	// Create tasks that add their index to the taskOrder slice
	tasks := []Task{
		func() error {
			mu.Lock()
			taskOrder = append(taskOrder, 1)
			mu.Unlock()
			return nil
		},
		func() error {
			mu.Lock()
			taskOrder = append(taskOrder, 2)
			mu.Unlock()
			return nil
		},
		func() error {
			mu.Lock()
			taskOrder = append(taskOrder, 3)
			mu.Unlock()
			return nil
		},
	}

	workerCount := 2
	_ = WorkerPool(tasks, workerCount)

	// Verify that all tasks were executed
	if len(taskOrder) != len(tasks) {
		t.Fatalf("expected %d tasks to be executed, but got %d", len(tasks), len(taskOrder))
	}
}

func TestWorkerPool_NoTasks(t *testing.T) {
	// Empty task list
	tasks := []Task{}

	workerCount := 3
	errors := WorkerPool(tasks, workerCount)

	// No errors should occur since no tasks exist
	if len(errors) != 0 {
		t.Fatalf("expected no errors, got %d: %v", len(errors), errors)
	}
}
