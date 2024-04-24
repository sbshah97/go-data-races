package main

import (
	"fmt"
	"sync"
	"time"
)

func CriticalSection(mutex sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	count++
}

var count int

func main() {
	var wg sync.WaitGroup

	mutex := sync.Mutex{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			CriticalSection(mutex)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			CriticalSection(mutex)
		}
	}()

	wg.Wait()
	time.Sleep(time.Second * 5)
	fmt.Println("Final counter value:", count)
}