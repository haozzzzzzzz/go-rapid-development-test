#!/usr/bin/env bash
export GOROOT=/usr/local/go
export GOPATH=/Users/hao/Documents/Projects/Github/go-rapid-development-test

if [ -e api ]
then
    rm api
fi

# api project
go build -o api ${GOPATH}/src/github.com/haozzzzzzzz/go-rapid-development/tools/api/main.go

