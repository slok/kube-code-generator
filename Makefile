

IMAGE := slok/kube-code-generator:latest

default: build

.PHONY: build
build:
	docker build -t $(IMAGE) .
