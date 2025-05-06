package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2) // number of go routines
	go firstTask()
	go secondTask()
	wg.Wait() // wait untill all go routines are done
}

func firstTask() {
	fmt.Println("first task")
	time.Sleep(3 * time.Second)
	wg.Done() // done
}

func secondTask() {
	fmt.Println("second task")
	time.Sleep(2 * time.Second)
	wg.Done() // done
}
