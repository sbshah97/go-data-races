package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func ProcessJob(job *Job) {
	// Simulate expensive task
	time.Sleep(time.Second*time.Duration(job.ID))
	fmt.Printf("Processing job with ID: %d\n", job.ID)
}

func main() {
	jobs := []Job{{ID: 1}, {ID: 2}, {ID: 3}}

	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	ProcessJob(job) // Here job is passed as an free variable argument
		// }()
		go func(job Job) {
			defer wg.Done()
			ProcessJob(&job) // Here job is passed as an free variable argument
		}(job)
	}
	wg.Wait()
}