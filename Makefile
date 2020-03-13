APP_NAME         := poker
GIT_REPO         := github.com/herry13/poker-go
PKGS             := $(shell go list ./... | grep -vF /vendor/)
RELEASE_DIR      := ./release
GO_BUILD_OPTS    := -a -installsuffix cgo

GOLANGCI_LINT     := golangci-lint-1.23.7-linux-amd64
GOLANGCI_LINT_URL := https://github.com/golangci/golangci-lint/releases/download/v1.23.7/$(GOLANGCI_LINT).tar.gz

default: build validate test

.tools/$(GOLANGCI_LINT)/golangci-lint:
	mkdir -p .tools && cd .tools && curl -sL $(GOLANGCI_LINT_URL) | tar xvz

build:
	go fmt ./...
	go build $(GO_BUILD_OPTS) -o $(RELEASE_DIR)/$(APP_NAME) $(GIT_REPO)/cmd/$(APP_NAME)

validate: .tools/$(GOLANGCI_LINT)/golangci-lint
	go vet $(PKGS)
	./.tools/$(GOLANGCI_LINT)/golangci-lint run

test:
	go test -race ./...

clean:
	rm -r $(RELEASE_DIR)

.PHONY: build validate test clean

