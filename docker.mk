.PHONY: help setup dependencies clean

NAME := tamakiii-sandbox/hello-golang-json

help:
	@cat $(firstword $(MAKEFILE_LIST))

setup: \
	dependencies

dependencies:
	type docker

build: Dockerfile
	docker build -t $(NAME) .

bash: build
	docker run -it --rm -v $(PWD):/go -w /go -e GOPATH=/go $(NAME) $@

clean:
