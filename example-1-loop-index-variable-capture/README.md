# Go Concurrency: Job Processing with Goroutines
This project demonstrates how to use goroutines in Go for concurrent job processing. It includes a simple example of a data race condition and how to avoid it.

# Code Overview

Go’s design choice to transparently capture free variables by reference in goroutines is a recipe for data races

The code defines a Job struct with an ID field. The ProcessJob function simulates an expensive task by sleeping for a second and then printing the job ID.

In the main function, we create a slice of Job instances and iterate over them. For each job, we start a goroutine to process the job concurrently.

# How to Run
To run the code, use the go run command:
```
go run main.go
```

## Note
This code is intended as a demonstration of Go's concurrency features and the potential for data races when using goroutines. Always be aware of the potential for data races when using goroutines and ensure that you use appropriate synchronization mechanisms, such as the sync package's WaitGroup, to prevent them.