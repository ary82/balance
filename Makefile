all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

run: build
	@echo "Running..."
	@go run cmd/api/main.go

test:
	@echo "Testing..."
	@go test -v ./...

clean:
	@echo "Cleaning Go server..."
	@rm -f main
	@echo "Cleaning Python server..."
	@rm -rf ./classification/__pycache__
	@rm -rf ./classification/.venv

py-server:
	@( \
		 cd ./classification; \
		 echo "making python venv"; \
		 python -m venv .venv; \
		 echo "sourcing python venv"; \
		 source ./.venv/bin/activate; \
		 echo "downloading dependencies"; \
		 pip install -r ./requirements.txt; \
		 echo "starting py-server"; \
		 python ./main.py; \
	)

# Live Reload
watch:
	@${HOME}/go/bin/air

.PHONY: all build run test clean py-server
