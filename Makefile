.PHONY: help setup dependencies build test clean

help:
	@cat $(firstword $(MAKEFILE_LIST))

setup: \
	dependencies

dependencies:
	type go

build:
	go build

test: build
	go test

clean:
