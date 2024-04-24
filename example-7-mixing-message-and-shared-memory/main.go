package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Future struct {
	sync.Mutex
	response string
	err      error
	ch       chan int
}

// func (f *Future) Start() {
// 	go func() {
// 		time.Sleep(2 * time.Second) // Simulate some work
// 		f.Lock()
// 		f.response = "Hello, world!"
// 		f.Unlock()
// 		f.ch <- 1 // Signal that the work is done
// 	}()
// }

// func (f *Future) Wait(ctx context.Context) error {
// 	select {
// 	case <-f.ch: // Wait for the work to be done
// 		f.Lock()
// 		defer f.Unlock()
// 		fmt.Println(f.response)
// 		return f.err
// 	case <-ctx.Done(): // Or for the context to be cancelled
// 		f.Lock()
// 		defer f.Unlock()
// 		f.err = errors.New("cancelled")
// 		return f.err
// 	}
// }

func (f *Future) Start() {
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work
		f.response = "Hello, world!"
		f.err = errors.New("execution error")
		f.ch <- 1 // Signal that the work is done
	}()
}

// Context cancel racing with work being done and overwriting the value / error
func (f *Future) Wait(ctx context.Context) error {
	select {
	case <-f.ch: // Wait for the work to be done
		return nil
	case <-ctx.Done(): // Or for the context to be cancelled
		f.err = errors.New("cancelled")
		return f.err
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	f := &Future{ch: make(chan int)}
	f.Start()
	err := f.Wait(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.response)
	}
}

// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type Future struct {
// 	sync.Mutex
// 	response string
// 	err      error
// 	ch       chan struct{}
// }

// func (f *Future) Start() {
// 	go func() {
// 		time.Sleep(2 * time.Second) // Simulate some work
// 		f.Lock()
// 		f.response = "Hello, world!"
// 		f.Unlock()
// 		close(f.ch) // Signal that the work is done
// 	}()
// }

// func (f *Future) Wait(ctx context.Context) error {
// 	select {
// 	case <-f.ch: // Wait for the work to be done
// 		f.Lock()
// 		defer f.Unlock()
// 		fmt.Println(f.response)
// 		return f.err
// 	case <-ctx.Done(): // Or for the context to be cancelled
// 		f.Lock()
// 		defer f.Unlock()
// 		f.err = errors.New("cancelled")
// 		return f.err
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	f := &Future{ch: make(chan struct{})}
// 	f.Start()
// 	err := f.Wait(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }