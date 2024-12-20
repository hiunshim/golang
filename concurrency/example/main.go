package main

import (
	"errors"
	"fmt"
	"sync"
)

// Task is a function that returns an error
type Task func() error

// WorkerPool processes tasks concurrently using a fixed number of workers.
func WorkerPool(tasks []Task, workerCount int) []error {
	var wg sync.WaitGroup
	taskCh := make(chan Task)
	errCh := make(chan error, len(tasks)) // Buffered to hold all potential errors

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskCh {
				if err := task(); err != nil {
					errCh <- err
				}
			}
		}()
	}

	// Send tasks to the taskCh
	go func() {
		for _, task := range tasks {
			taskCh <- task
		}
		close(taskCh)
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(errCh)

	// Collect errors
	var errors []error
	for err := range errCh {
		errors = append(errors, err)
	}

	return errors
}

func main() {
	// Example tasks
	tasks := []Task{
		func() error { return nil },
		func() error { return errors.New("Task 2 failed") },
		func() error { return nil },
		func() error { return errors.New("Task 4 failed") },
	}

	// Run WorkerPool
	workerCount := 3
	errors := WorkerPool(tasks, workerCount)

	// Print results
	if len(errors) == 0 {
		fmt.Println("All tasks completed successfully")
	} else {
		fmt.Println("Errors:")
		for _, err := range errors {
			fmt.Println(err)
		}
	}
}
