// This is an example to demo that channel could be shared with
// many senders. This could be useful when worker need to process
// from a diverse source of inputs
//
// There are multiple notifiers in this example, and processor receives from all of
// them and print their messages.

package main

import (
	"fmt"
	"log"
	"time"
)

func notifier(id int, notifyChan chan<- string) {
	ticker := time.NewTicker(time.Second)

	for t := range ticker.C {
		notifyChan <- fmt.Sprintf("Notification from %d: at %s", id, t)
	}
}

func processor(notifyChan <-chan string) {
	for n := range notifyChan {
		log.Printf("Received notification: %s", n)
	}
}

func main() {
	notifyChan := make(chan string, 3)

	for i := 0; i < 5; i++ {
		go notifier(i, notifyChan)
	}

	processor(notifyChan)
}
