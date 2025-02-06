package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Worker Interface demonstration
type Worker interface {
	DoWork() string
}

// BaseWorker Struct with embedded type
type BaseWorker struct {
	name string
}

// CustomWorker implementing Worker interface
type CustomWorker struct {
	BaseWorker
	jobType string
}

func (w *CustomWorker) DoWork() string {
	return fmt.Sprintf("Worker %s doing %s job", w.name, w.jobType)
}

// CustomError Error handling demonstration
type CustomError struct {
	code    int
	message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

// Channel worker function
func processWork(ctx context.Context, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// Simulate work
			time.Sleep(100 * time.Millisecond)
			results <- job * 2
		case <-ctx.Done():
			fmt.Println("Worker cancelled")
			return
		}
	}
}

func main() {
	// 1. Goroutines and channels demonstration
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go processWork(ctx, jobs, results, &wg)
	}

	// Send jobs
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Collect results using select
	go func() {
		wg.Wait()
		close(results)
	}()

	// 2. Interface usage demonstration
	worker := &CustomWorker{
		BaseWorker: BaseWorker{name: "John"},
		jobType:    "processing",
	}
	fmt.Println(worker.DoWork())

	// 3. Error handling
	if err := doSomething(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	}

	// 4. Map with mutex for concurrent access
	cache := struct {
		sync.RWMutex
		data map[string]int
	}{
		data: make(map[string]int),
	}

	// Concurrent map access
	go func() {
		cache.Lock()
		cache.data["key"] = 42
		cache.Unlock()
	}()

	// Read from map
	cache.RLock()
	if val, ok := cache.data["key"]; ok {
		fmt.Printf("Value: %d\n", val)
	}
	cache.RUnlock()

	// 5. Collect results
	for result := range results {
		fmt.Printf("Got result: %d\n", result)
	}
}

func doSomething() error {
	return &CustomError{code: 500, message: "something went wrong"}
}
