all: build

build:
	@echo "building go binary..."
	@go build -o main cmd/app/main.go
	@echo "building tailwind css..."
	@pnpm run tailwind-build

run: build
	@echo "running..."
	@./main

clean:
	@echo "cleaning Go server..."
	@rm -f main
	@rm -f ./static/styles/output.css
	@echo "cleaning Python server..."
	@rm -rf ./proto/__pycache__
	@rm -rf ./.venv

py-server:
	@( \
		 echo "making python venv"; \
		 python -m venv .venv; \
		 echo "sourcing python venv"; \
		 source ./.venv/bin/activate; \
		 echo "downloading dependencies"; \
		 pip install -r ./requirements.txt; \
		 echo "starting py-server"; \
		 python ./cmd/classification_grpc/main.py; \
	)

# Live Reload
watch:
	@${HOME}/go/bin/air

.PHONY: all build run clean py-server
