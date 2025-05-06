// https://claude.ai/public/artifacts/3d595e28-fa44-45ec-a0bb-635bc8f2cbb0
package main

import (
	"fmt"
	"sync"
	"time"
)

// task performs work identified by id, sends a completion message to taskStatusChan,
// and signals completion through WaitGroup
func task(id int, taskStatusChan chan string, wg *sync.WaitGroup) {
	// Ensure we mark this task as done when the function exits,
	// even if there's an error/panic later
	defer wg.Done()

	// Log that the task is starting
	fmt.Printf("Task %d starting..\n", id)

	// Simulate work - each task takes a different amount of time
	// proportional to its ID
	time.Sleep(time.Duration(id) * time.Second)

	// Send completion message to the channel
	// NOTE: This could block if no one is receiving from the channel
	taskStatusChan <- fmt.Sprintf("Task %d done!", id)
}

// closeWhenDone waits for all tasks to complete then closes the channel
// This is a named function alternative to the anonymous function approach
func closeWhenDone(wg *sync.WaitGroup, taskStatusChan chan string) {
	// Wait for all tasks to signal completion
	wg.Wait()

	// Close the channel to signal no more messages will be sent
	// This allows any range loops to terminate naturally
	close(taskStatusChan)
}

func main() {
	// Create an unbuffered channel for task status messages
	taskStatusChan := make(chan string)

	// Create a WaitGroup to track completion of all tasks
	var wg sync.WaitGroup

	// Start 3 tasks concurrently
	for i := 1; i <= 3; i++ {
		// Increment counter BEFORE starting the goroutine
		wg.Add(1)

		// Launch the task as a goroutine
		go task(i, taskStatusChan, &wg)
	}

	// Start a dedicated goroutine to close the channel when all tasks complete
	// This is CRITICAL to prevent deadlock:
	// - If we tried to Wait() in the main goroutine, we'd block and never read from the channel
	// - If we never closed the channel, the range loop below would block forever
	go func() {
		wg.Wait()
		close(taskStatusChan)
	}()

	// Alternative approach using a named function instead of an anonymous function
	// This is functionally identical to the anonymous function above
	// go closeWhenDone(&wg, taskStatusChan)

	// Read and print all messages from the channel until it's closed
	// The for-range loop automatically exits when the channel is closed
	for taskStatus := range taskStatusChan {
		fmt.Println(taskStatus)
	}

	// This line executes only after all tasks have completed and
	// the channel has been closed
	fmt.Println("All task done")
}
