.DEFAULT_GOAL := build

.PHONY:clean fmt vet build
clean:
	go clean
	rm gozone

fmt: clean
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o gozone ./cmd/gozone
