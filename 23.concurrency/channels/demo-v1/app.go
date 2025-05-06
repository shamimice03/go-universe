package main

import (
	"fmt"
	"time"
)

func main() {
	firstChan := make(chan bool)
	secondChan := make(chan string)

	go validation(firstChan)
	go calculation(firstChan, secondChan)

	fmt.Println(<-secondChan)

}

func validation(firstChan chan bool) {
	fmt.Println("Validation")
	time.Sleep(3 * time.Second)
	firstChan <- true
}

func calculation(firstChan chan bool, secondChan chan string) {
	fmt.Println("Calculation")
	time.Sleep(2 * time.Second)

	// waiting for firstChan to finish
	if <-firstChan {
		secondChan <- "Done"
	}
}
