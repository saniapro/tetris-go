.PHONY: build build-stripped lint fmt vet test clean help

# Variables
BINARY_NAME=tetris
BINARY_PATH=./bin/$(BINARY_NAME)
CMD_PATH=./cmd/tetris/main.go
GO=go

help:
	@echo "Available targets:"
	@echo "  build           - Build the binary (debug)"
	@echo "  build-stripped  - Build stripped binary (no symbols/debug info)"
	@echo "  lint            - Run golangci-lint (if available)"
	@echo "  fmt             - Format code with gofmt"
	@echo "  vet             - Run go vet for static analysis"
	@echo "  test            - Run tests"
	@echo "  clean           - Remove build artifacts"

build:
	@mkdir -p bin
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BINARY_PATH) $(CMD_PATH)
	@echo "✓ Build complete: $(BINARY_PATH)"

build-stripped:
	@mkdir -p bin
	@echo "Building stripped $(BINARY_NAME)..."
	$(GO) build -ldflags="-s -w" -o $(BINARY_PATH) $(CMD_PATH)
	@echo "✓ Stripped build complete: $(BINARY_PATH)"

lint:
	@echo "Running golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "⚠ golangci-lint not installed. Install with: https://golangci-lint.run/usage/install/"; \
	fi

fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...
	@echo "✓ Format complete"

vet:
	@echo "Running go vet..."
	$(GO) vet ./...
	@echo "✓ Vet complete"

test:
	@echo "Running tests..."
	$(GO) test -v ./...

clean:
	@echo "Cleaning..."
	rm -rf bin/
	$(GO) clean
	@echo "✓ Clean complete"
