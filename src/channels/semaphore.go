// This is an example to use the blocking nature of
// channel as a semaphore. Once all elements in a channel
// is taken, getting items from the channel will block

package main

import (
	"log"
	"sync"
	"time"
)

func workSemaphore(i int) {
	time.Sleep(time.Second)
	log.Printf("Processing work %d", i)
}

func main() {
	// init semaphore by adding items in the channel
	sem := make(chan struct{}, 4)
	for i := 0; i < 4; i++ {
		sem <- struct{}{}
	}

	wg := &sync.WaitGroup{}

	// although 10 workers are launched, the semaphore implementation guarantees
	// there are max 4 workers running in parallel
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			// each time a worker starts, take one out of the
			// channel and start processing
			// Once all items are taken out, it'll block and wait
			// for the previous workers to finish
			<-sem
			defer func() {
				// remember to give back to the channel once job is done
				sem <- struct{}{}
			}()
			defer func() {
				wg.Done()
			}()

			workSemaphore(i)
		}(i)
	}

	wg.Wait()
}
