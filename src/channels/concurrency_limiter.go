// This implements a simple concept of a rate limiter in Golang with channels.
//
// It'll come handy when you want to limit the number of concurrent
// accesses to a shared resources

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ch chan struct{}) {
	time.Sleep(time.Second)

	fmt.Printf("Running worker %d\n", id)

}

func main() {
	work := make([]int, 100)
	for i := 0; i < 100; i++ {
		work[i] = i
	}

	// limit to 10 workers
	ch := make(chan struct{}, 10)

	fmt.Println("Work starting...")

	start := time.Now()
	wg := &sync.WaitGroup{}
	for _, w := range work {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			// holds the limit when starting, and takes it back
			// once work is done
			ch <- struct{}{}
			defer func() {
				<-ch
			}()

			worker(w, ch)
		}(w)
	}
	wg.Wait()

	end := time.Now()
	fmt.Printf("All work done in %v\n", end.Sub(start))
}
