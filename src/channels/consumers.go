// This example demonstrates the producer-consumer pattern in Go, where one
// goroutine (the producer) generates data and multiple other goroutines (the
// consumers) process it concurrently.
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// producer creates a new channel and launches a goroutine to populate it with
// data. In this case, it sends the integers 0 through 15 to the channel. Once
// all the data has been sent, the producer closes the channel to signal that
// no more data will be sent.
func producer() chan int {
	result := make(chan int)

	go func() {
		for i := 0; i < 16; i++ {
			fmt.Printf("Producing item %d\n", i)
			result <- i
		}
		close(result)
	}()

	return result
}

// consumer receives a channel of integers and processes each item from the
// channel. In this case, it simulates work by sleeping for one second for each
// item and then logging a message. The consumer continues to process items
// until the channel is closed and all values have been received.
func consumer(inputChan chan int) {
	for i := range inputChan {
		time.Sleep(time.Second)
		log.Printf("Consuming item %d", i)
	}
}

// main is the entry point of the program. It creates a channel of data using
// the producer function, and then starts four consumer goroutines to process
// the data from the channel. A WaitGroup is used to ensure that the main
// function waits for all the consumer goroutines to finish their work before
// exiting.
func main() {
	inputChan := producer()

	log.Printf("Start working...")

	start := time.Now()

	wg := &sync.WaitGroup{}

	// spawns 4 consumers for the work
	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
			}()

			consumer(inputChan)
		}()
	}

	wg.Wait()
	end := time.Now()

	log.Printf("All work done in %v", end.Sub(start))
}