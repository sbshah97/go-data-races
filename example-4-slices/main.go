package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

func ProcessUuids(uuids []string) {
	// Simulate expensive task
	// time.Sleep(time.Second*time.Duration(job.ID))
	// fmt.Printf("Processing job with ID: %d\n", job.ID)
	var myResults []string
	var mutex sync.Mutex
	safeAppend := func(result string) {
		mutex.Lock()
		myResults = append(myResults, result)
		mutex.Unlock()
	}

	for _, uuid := range uuids {
		go func(id string, results []string) {
			// Do something with uuid
			res := id + " processed"
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))

			safeAppend(res)
		}(uuid, myResults)
	}

	time.Sleep(time.Second * 15)
	fmt.Println("Count of myResults: ", len(myResults))
}

func main() {
	var uuids []string
	// Generate a slice of uuids
	for i := 0; i < 10000; i++ {
		uuids = append(uuids, uuid.New().String())
	}
	ProcessUuids(uuids)
}
