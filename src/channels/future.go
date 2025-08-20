// This program demonstrates a basic implementation of a "Future" concurrency
// pattern in Go. (like Python async/await pattern).
// A future represents a value that may not be available yet, but
// will be at some point in the future.
// This implementation achieves this by saves the result in a channel.
// The consumer waits on this channel when "awaiting" the result from this Future.

package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Future is a struct that holds a channel for receiving a future result, and an
// error field to store any potential errors that occurred during the asynchronous
// operation.
type Future struct {
	Chan chan interface{}
	Err  error
}

// NewFuture creates and returns a new Future object. The channel within the
// Future is initialized, ready to receive a value.
func NewFuture() Future {
	return Future{
		Chan: make(chan interface{}),
	}
}

// Get is a method on the Future struct that blocks until a result is available
// on the Future's channel. If the Future has an error, it returns the error
// immediately. Once the result is received, the channel is closed to prevent
// further reads, and the result is returned.
func (fut *Future) Get() (interface{}, error) {
	if fut.Err != nil {
		return nil, fut.Err
	}

	result, ok := <-fut.Chan
	if !ok {
		return nil, errors.New("Future already closed")
	}
	close(fut.Chan)
	return result, nil
}

// work is an example of a function that returns a Future. It starts a new
// goroutine to perform some work (in this case, sleeping for a second and then
// creating a string). The result of the work is sent to the Future's channel.
func work(i int) Future {
	result := NewFuture()
	go func() {
		time.Sleep(1 * time.Second)

		result.Chan <- fmt.Sprintf("Hello World %d", i)
	}()
	return result
}

// main is the entry point of the program. It demonstrates how to use the Future
// implementation by launching several asynchronous 'work' operations, and then
// collecting their results using the 'Get' method on the returned Future objects.
func main() {
	futs := []Future{}

	log.Printf("Starting work async...")
	// Launch 5 asynchronous 'work' operations. Each call to 'work' returns a
	// Future immediately.
	for i := 0; i < 5; i++ {
		futs = append(futs, work(i))
	}

	log.Printf("All work dispatched, now collecting results...")
	// Iterate over the slice of Futures and call 'Get' on each one to retrieve
	// the result. This will block until the result is available.
	for i, fut := range futs {
		if val, err := fut.Get(); err != nil {
			log.Printf("ERRO: getting result %d: %s", i, err)
		} else {
			result := val.(string)

			log.Printf("INFO: Getting result %d: %s", i, result)
		}
	}
	log.Print("All work finished...")

	// This demonstrates that reading from a closed future returns an error.
	log.Print("Below is an expected error:")
	_, err := futs[0].Get()
	if err != nil {
		log.Printf("Error getting closed future: %s", err)
	}
}