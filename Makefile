.PHONY: help setup dependencies test clean

help:
	@cat $(firstword $(MAKEFILE_LIST))

setup: \
	dependencies

dependencies:
	type go

test:
	go test src/main/main_test.go

clean:
