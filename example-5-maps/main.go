package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	m := make(map[int]int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			m[i] = i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			fmt.Println(m[i])
		}
	}()

	wg.Wait()
}