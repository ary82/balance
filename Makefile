all: build

build:
	@echo "Building go binary..."
	@go build -o main cmd/api/main.go
	@echo "Building tailwind css..."
	@pnpm run tailwind-build

run: build
	@echo "Running..."
	@./main

clean:
	@echo "Cleaning Go server..."
	@rm -f main
	@rm -f ./static/styles/output.css
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

.PHONY: all build run clean py-server
