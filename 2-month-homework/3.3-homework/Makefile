# Define variables
APP_NAME := myapp
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := "-X main.Version=$(VERSION)"

# Define commands
.PHONY: all build clean test run lint fmt vet deps install update-deps

run:
	@echo "Running the application..."
	go run $(GO_FILES)