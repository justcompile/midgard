#!/usr/bin/env bash

set -e

LINTER=${GOPATH}/bin/golangci-lint

lint() {
    pkg="$1"
    ${LINTER} run -D structcheck -E interfacer -E maligned -E prealloc -E depguard -E gocyclo -E gosec -E nakedret -E typecheck -e S1000 -e "field.* is unused" --deadline 180s "$pkg"
}

lint ./...
