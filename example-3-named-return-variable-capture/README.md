# Go Concurrency: Data Race due to Named Return Variable
This repository contains a Go code example that demonstrates a data race condition due to named return variable. In Go, a function can name its return variable, and this named return variable is accessible throughout the function. However, when this feature is mixed with a goroutine, the named return variable gets captured by reference in the closure, which can lead to a data race.

## Code Overview
The code simulates a scenario where a function with a named return variable is called. The function checks the value of a global variable value. If value is 0, the function returns immediately. Otherwise, it starts a goroutine that assigns the address of the named return variable to value, and then returns 20 after a delay of 2 seconds.

The main function sets value to different values and calls the function, printing the return value and the value pointed to by value.

## How to Run
To run the code, use the go run command:
```
go run main.go
```

## Note
This code is intended as a demonstration of a potential data race condition in Go when using named return variables inside a goroutine. Always be aware of the potential for data races when using goroutines and ensure that you use appropriate synchronization mechanisms, such as the sync package's WaitGroup, to prevent them.

In this code, the NamedReturn function has a named return variable result. The function checks the value of a global variable value. If value is 0, the function returns immediately. Otherwise, it starts a goroutine that assigns the address of result to value, and then returns 20 after a delay of 2 seconds. This can lead to a data race because the goroutine and the return statement in the NamedReturn function are both accessing the result variable concurrently.