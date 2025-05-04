package main

import (
	"fmt"
	"time"
)

func first(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func second(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func slow(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", phrase)
	doneChan <- true
	close(doneChan) // closing here since, i know this function will take longest.
}

func main() {

	done := make(chan bool)
	go first("from first function", done)
	go second("from second function", done)
	go slow("from slow function", done)
	// <-done
	// <-done
	// <-done

	for range done {

	}
}
