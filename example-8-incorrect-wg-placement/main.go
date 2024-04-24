package main

import (
	"fmt"
	"sync"
)

func processItem(id int, results []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate some work with a random result
	results[id] = id * 2
}

func WaitGrpExample(itemIds []int) {
	var wg sync.WaitGroup
	results := make([]int, len(itemIds))

	for i := range itemIds {
		go func(i int) {
			wg.Add(1)
			processItem(i, results, &wg)
		}(i)
	}

	wg.Wait()

	successCount := 0
	for _, result := range results {
		if result > 0 {
			successCount++
		}
	}

	fmt.Printf("Number of successful processings: %d\n", successCount)
}

func main() {
	itemIds := []int{1, 2, 3, 4, 5}
	WaitGrpExample(itemIds)
}