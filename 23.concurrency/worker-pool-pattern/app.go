package main

import (
	"fmt"
	"net/http"
	"sync"
)

var api_url_list []string

func main() {
	api_url_list = []string{
		"https://github.com/",
		"https://medium.com/",
		"https://dev.to/",
		"https://linkedin.com/",
		"https://aws.amazon.com/",
		"https://render.com/",
		"https://udemy.com/",
		"https://stackoverflow.com/",
		// imagine 1000s of url/api endpoint
	}

	tasks := make(chan string, len(api_url_list))
	results := make(chan string, len(api_url_list))

	// Number of worker goroutines to use
	numWorkers := 3

	// Start the worker pool
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, tasks, results)
	}

	// Send all URLs to the task channel
	for _, url := range api_url_list {
		tasks <- url
	}
	close(tasks) // Close the tasks channel to signal no more tasks

	// Wait for all workers to finish in a separate goroutine
	go func() {
		wg.Wait()
		close(results) // Close results channel when all workers are done
	}()

	// Collect and display results
	for result := range results {
		fmt.Println(result)
	}
}

func worker(id int, wg *sync.WaitGroup, tasks <-chan string, results chan<- string) {
	defer wg.Done()

	// Process URLs from the tasks channel
	for url := range tasks {
		fmt.Printf("Worker %d: Checking %s\n", id, url)

		response, err := http.Get(url)

		if err != nil {
			results <- fmt.Sprintf("Worker %d: URL %s is not responding: %v", id, url, err)
			continue
		}

		results <- fmt.Sprintf("Worker %d: Status of %s is %s", id, url, response.Status)

		// Don't forget to close the response body to prevent resource leaks
		response.Body.Close()
	}
}
