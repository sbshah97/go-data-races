run-example-1:
	go run -race example-1-loop-index-variable-capture/main.go

run-example-2:
	go run -race example-2-err-variable-capture/main.go

run-example-3:
	go run -race example-3-named-return-variable-capture/main.go

run-example-4:
	go run -race example-4-slices/main.go

run-example-5:
	go run -race example-5-maps/main.go

run-example-6:
	go run -race example-6-pass-by-value/main.go

run-example-7:
	go run -race example-7-mixing-message-and-shared-memory/main.go

run-example-8:
	go run -race example-8-incorrect-wg-placement/main.go