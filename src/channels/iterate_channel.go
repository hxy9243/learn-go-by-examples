// This is an example of iterating through a channel to get results
// The sender will block on the channel if there's no buffer or buffer is full.
//
// How much time will it take to finish iterating through all the results?

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func sender(results chan<- string) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			time.Sleep(time.Second)
			results <- fmt.Sprintf("result %d", i)
		}(i)
	}

	// after all the results are sent, we can now
	// safely close the channel
	wg.Wait()
	close(results)
}

func receiver(results <-chan string) {
	for r := range results {
		log.Printf("Getting result %s\n", r)
	}
}

func main() {
	log.Printf("Starting program...")

	// creating a blocked channel
	// note: in certain cases, you can create a buffered channel
	// so that the sender goroutines can quit early
	results := make(chan string)

	start := time.Now()

	go sender(results)

	receiver(results)

	end := time.Now()

	fmt.Printf("All receiving finished in %v\n", end.Sub(start))
}
