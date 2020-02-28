# Poker Go

[![CircleCI](https://circleci.com/gh/herry13/poker-go.svg?style=svg)](https://circleci.com/gh/herry13/poker-go)

## Overview

Poker game server implemented in Go language

## Dependencies

- [Go](https://golang.org) (>= 1.13)

## Build

To build this project:

```
make build
```

## Validate source code

Source code validation is performed using `go vet` and `golint`. To validate the source code:

```
make validate
```

## Running the tests

```
make test
```

## Adding new dependencies

Dependencies required by `poker_go` should be added using `dep` command.
For example, to add a dependency on `gopkg.in/yaml.v2` with version `2.2.2`:

```
dep ensure -add gopkg.in/yaml.v2@2.2.2
```

