# Go Channels Basic Examples

This directory contains a collection of examples demonstrating various use cases for Go channels. Each example is a self-contained program that can be run to observe its behavior.

## Quick Start

Make sure you have Go installed on your system.

Run:

```
go run consumer.go
```

## Examples

- `consumers.go`: Demonstrates the **producer-consumer** pattern. A single producer goroutine generates data and sends it to a channel, while multiple consumer goroutines concurrently receive and process the data from the channel.
- `context.go`: Shows how to use the `context` package to **cancel a running goroutine**.
- `future.go`: Implements a simple **"future"**-like interface using channels. A function can return a `Future` object, which is a wrapper around a channel. The caller can then call the `Get()` method on the `Future` to block and wait for the result of the asynchronous operation.
- `get_result.go`: A basic example of how to receive results from asynchronous functions using channels. Multiple worker goroutines are launched, and they send their results to a shared channel. The main goroutine then receives the results from the channel.
- `iterate_channel.go`: Shows how to iterate over a channel using a `for...range` loop. A sender goroutine sends multiple values to a channel and then closes it. The receiver goroutine then uses a `for...range` loop to receive all the values from the channel until it is closed.
- `message_broker.go`: Demonstrates a simple **message broker** pattern. The main goroutine creates a set of worker goroutines and a corresponding set of channels. It then sends a "ready" signal to each worker on its respective channel to start its work.
- `multiple_source.go`: Illustrates how to receive values from **multiple sources** into a single channel. Multiple notifier goroutines are created, and they all send notifications to the same channel. A single processor goroutine then receives and processes all the notifications from that channel.
- `semaphore.go`: Shows how to use a buffered channel as a **semaphore** to limit the number of concurrent goroutines. A buffered channel is pre-filled with a certain number of empty structs. To start a new task, a goroutine must first receive a value from the channel. When the task is complete, it sends a value back to the channel, allowing another goroutine to proceed.
- `signal.go`: Demonstrates how to handle **system signals** (e.g., `SIGINT`, `SIGTERM`) using channels. The `os/signal` package is used to create a channel that receives notifications of incoming signals. A worker goroutine can then listen on this channel and take appropriate action when a signal is received.
-`timer.go`: Shows how to use `time.Ticker` and `time.Timer` to perform periodic tasks and handle timeouts. A worker goroutine uses a `time.Ticker` to perform a task at regular intervals. A `time.Timer` is used to set a timeout for the entire operation.