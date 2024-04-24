package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	var err error
	err = errors.New("serial error initial function") // This is the error we want to capture
	if err != nil {
		fmt.Printf("Value of error: %v\n", err)
	}

	// Here we start a goroutine that will run concurrently with the rest of the function.
	// Inside this goroutine, we assign a new error to the err variable if the id is even.
	// This err variable is the same one defined in the outer function, so it's captured by reference in this closure.
	// This means that the goroutine and the outer function are both reading and writing to the err variable at the same time,
	// which can lead to a data race. A data race occurs when two or more goroutines access the same variable concurrently,
	// and at least one of the accesses is a write, and the goroutines are not using any kind of synchronization like a mutex
	// to control their access to the variable.
	go func() {
		for id := 0; id < 10; id++ {
			// defer wg.Done()
			// Simulate a function that may return an error
			time.Sleep(time.Second * time.Duration(id))

			if id%2 == 0 {
				err = errors.New("goroutine error")
				if err != nil {
					fmt.Printf("Value of Error in goroutine %d: %v\n", id, err)
				}
			}
		}
	}()

	err = errors.New("serial error end of function")
	if err != nil {
		time.Sleep(time.Second * 2)
		fmt.Printf("Value of error at end of function: %v\n", err)
	}

}