APP_NAME         := poker_go
GIT_REPO         := github.com/herry13/poker-go
PKGS             := $(shell go list ./... | grep -vF /vendor/)
RELEASE_DIR      := ./release
GO_BUILD_OPTS    := -a -installsuffix cgo


default: build validate test

build:
	go fmt ./...
	go build $(GO_BUILD_OPTS) -o $(RELEASE_DIR)/$(APP_NAME) $(GIT_REPO)/cmd/$(APP_NAME)

validate:
	go vet $(PKGS)
	golint $(PKGS)

test:
	go test -race ./...

clean:
	rm -r $(RELEASE_DIR)

.PHONY: build validate test clean

