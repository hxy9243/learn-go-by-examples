// This is as an example of receiving result from async functions
// using golang channel.
// Instead of packaging the result and returns in a Future, we can
// send a channel to the calling function, and reads the result back.

package main

import (
	"fmt"
	"log"
	"time"
)

// Result is the result type from the work functions
type Result struct {
	Val int
	Err error
}

// simulate very simple workload: sleep 1 second, return id + 1
func worker(id int, resultChan chan Result) {
	time.Sleep(2 * time.Second)

	resultChan <- Result{Val: id + 1}
}

func main() {
	resultChan := make(chan Result)

	log.Print("work starting...")
	for i := 0; i < 5; i++ {
		// carefully: id here needs to be passed in as parameter
		// to async function
		// goroutines are started async, and may read variables
		// in closure after the iteration is finished and index i is updated
		go func(i int) {
			worker(i, resultChan)
		}(i)
	}

	fmt.Println("Start to receiving results...")

	start := time.Now()
	var results []int
	for i := 0; i < 5; i++ {
		res := <-resultChan
		if res.Err != nil {
			log.Printf("Error: unable to get result for %d: %s", i, res.Err)
			continue
		}
		results = append(results, res.Val)
	}
	end := time.Now()

	log.Printf("Work finished after %v\n", end.Sub(start))
	log.Printf("Results: %v", results)
}
