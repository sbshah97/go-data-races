# Go Concurrency: Data Race with Map
This repository contains a Go code example that demonstrates a data race condition with a map. In Go, the built-in map type is not safe for concurrent use. If multiple goroutines access the same map concurrently, and at least one of them writes to the map, a data race occurs.

## Code Overview
The code simulates a scenario where two goroutines are accessing the same map. One goroutine is writing to the map, and the other is reading from it. Since Go's built-in maps are not safe for concurrent use, this will cause a data race.

## How to Run
To run the code, use the go run command with the -race flag to enable the data race detector:
```
go run -race main.go
```

If a data race is detected, a description of the race will be printed to the console.

## Note
This code is intended as a demonstration of a potential data race condition in Go when using maps. Always be aware of the potential for data races when using goroutines and ensure that you use appropriate synchronization mechanisms, such as the sync package's Mutex, to prevent them.

```
@sbshah97 ➜ /workspaces/codespaces-blank/example-5-data-races-in-maps $ go run -race main.go
0
0
==================
WARNING: DATA RACE
Write at 0x00c00007a090 by goroutine 7:
  runtime.mapassign_fast64()

Unable

```
@sbshah97 ➜ /workspaces/codespaces-blank/example-5-data-races-in-maps $ go run -race main.go
0
0
==================
WARNING: DATA RACE
Write at 0x00c00007a090 by goroutine 7:
  runtime.mapassign_fast64()
      /usr/local/go/src/runtime/map_fast64.go:93 +0x0
  main.main.func1()
      /workspaces/codespaces-blank/example-5-data-races-in-maps/main.go:16 +0xc6

Previous read at 0x00c00007a090 by goroutine 8:
  runtime.mapaccess1_fast64()
      /usr/local/go/src/runtime/map_fast64.go:13 +0x0
  main.main.func2()
      /workspaces/codespaces-blank/example-5-data-races-in-maps/main.go:23 +0xd0

Goroutine 7 (running) created at:
  main.main()
      /workspaces/codespaces-blank/example-5-data-races-in-maps/main.go:13 +0x10c

Goroutine 8 (running) created at:
  main.main()
      /workspaces/codespaces-blank/example-5-data-races-in-maps/main.go:20 +0x1bc
==================
fatal error: concurrent map read and map write

goroutine 7 [running]:
main.main.func2()
        /workspaces/codespaces-blank/example-5-data-races-in-maps/main.go:23 +0xd1
created by main.main in goroutine 1
        /workspaces/codespaces-blank/example-5-data-races-in-maps/main.go:20 +0x1bd

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc0000140e8?)
        /usr/local/go/src/runtime/sema.go:62 +0x25
sync.(*WaitGroup).Wait(0xc0000140e0)
        /usr/local/go/src/sync/waitgroup.go:116 +0xa5
main.main()
        /workspaces/codespaces-blank/example-5-data-races-in-maps/main.go:27 +0x1c7
exit status 2
```