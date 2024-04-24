# Go Concurrency: Incorrect WaitGroup Usage
This repository contains a Go code example that demonstrates a common mistake when using the sync.WaitGroup structure for group synchronization.

## Code Overview
The code defines two functions, processItem and WaitGrpExample.

processItem is a function that simulates some work by writing a result to a shared results slice. It takes an item id, a results slice, and a pointer to a sync.WaitGroup as arguments.

WaitGrpExample is a function that creates a goroutine for each item in a given slice to process the items concurrently. Each goroutine records its success or failure status in the results slice. However, wg.Add(1) is called inside the goroutine, which is not guaranteed to be executed before WaitGrpExample calls wg.Wait(). As a result, wg.Wait() can unblock prematurely, and WaitGrpExample can start reading from the results slice while some goroutines are still writing to it, causing a data race.

How to Run
To run the code, use the go run command:
```
go run -race main.go
```