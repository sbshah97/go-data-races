# Go Concurrency and Data Races
This repository contains code examples and discussions related to Uber's experience with Go concurrency and data races. Uber has adopted Go as a primary language for developing microservices, with a monorepo consisting of about 50 million lines of code and approximately 2,100 unique Go services.

## Go Concurrency
Go makes concurrency a first-class citizen, with the go keyword running function calls asynchronously. These asynchronous function calls, known as goroutines, are used by developers to hide latency. Goroutines can communicate data either via message passing (channels) or shared memory, with shared memory being the most commonly used means of data communication in Go.

## Data Races
Due to the ease of creating goroutines, Go programs typically expose significantly more concurrency than programs written in other languages. This higher concurrency also means a potential for more concurrency bugs. A data race is a concurrency bug that occurs when two or more goroutines access the same datum, at least one of them is a write, and there is no ordering between them. Data races are insidious bugs and must be avoided at all costs.

## Session Material can be found here:
* Slides: https://docs.google.com/presentation/d/1MxkngSPFGFB-ODqxF401KCqMPAZ1IgdhUgGmicFOwhg/edit?usp=sharing
* Running code examples: https://github.com/sbshah97/go-data-races
* Uber's Blog Post: https://www.uber.com/en-GB/blog/data-race-patterns-in-go/
* Uber's Paper: https://arxiv.org/abs/2204.00764Session Material can be found here:
Slides: https://docs.google.com/presentation/d/1MxkngSPFGFB-ODqxF401KCqMPAZ1IgdhUgGmicFOwhg/edit?usp=sharing