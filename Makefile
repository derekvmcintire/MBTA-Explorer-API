# Go binary name
BINARY_NAME=mbta-train-tracker-api

# Go-related variables
GO_CMD=go
GOFMT=go fmt
GOLINT=golangci-lint
LINTER_FLAGS=run

# Default target - build the API binary
.PHONY: all
all: build

# Build the Go project
.PHONY: build
build:
	$(GO_CMD) build -o $(BINARY_NAME) main.go

# Run the Go project locally
.PHONY: run
run:
	$(GO_CMD) run cmd/api/main.go

# Format Go code
.PHONY: fmt
fmt:
	$(GOFMT) ./...

# Run linter
.PHONY: lint
lint:
	$(GOLINT) $(LINTER_FLAGS)

# Run tests
.PHONY: test
test:
	$(GO_CMD) test ./...

# Clean up the binary
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

# Show help
.PHONY: help
help:
	@echo "Makefile for Go project"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build    - Build the Go binary"
	@echo "  run      - Run the Go application"
	@echo "  fmt      - Format the Go code"
	@echo "  lint     - Run linter on the code"
	@echo "  test     - Run Go tests"
	@echo "  clean    - Clean up the binary"
	@echo "  help     - Show this help message"

