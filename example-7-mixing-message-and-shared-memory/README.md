# Go Concurrency: Future Implementation with Channel Signalling
This repository contains a Go code example that demonstrates a Future implementation using channel signalling and shared memory.

## Code Overview
The code defines a Future struct that has a response string, an error, and a channel ch used for signalling. The Start method starts a goroutine that simulates some work, sets the response, and then sends a signal (an integer) on the channel to indicate that the work is done. The Wait method waits either for the work to be done or for a context to be cancelled. If the context is cancelled, it sets the error to "cancelled".

## How to Run
To run the code, use the go run command:
```
go run main.go
```

## Note
This code is intended as a demonstration of a Future implementation in Go using channel signalling and shared memory. Always be aware of the potential for data races when using goroutines and ensure that you use appropriate synchronization mechanisms, such as the sync package's Mutex, to prevent them.