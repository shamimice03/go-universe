package main

import (
	"fmt"
	"time"
)

func main() {
	// Create channels
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	// Sender goroutines
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message from channel 2"
	}()

	// Cancellation signal after 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	// Using select to handle multiple channels
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		case <-done:
			fmt.Println("Received done signal, exiting")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout - no message received in 500ms")
		}
	}
}
