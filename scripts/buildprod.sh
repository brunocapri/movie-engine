#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o movie-engine ./cmd/server/main.go
