all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

run:
	@go run cmd/api/main.go

test:
	@echo "Testing..."
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@${HOME}/go/bin/air

.PHONY: all build run test clean
