

IMAGE := slok/kube-code-generator:latest

default: build

.PHONY: build
build:
	docker buildx build --platform linux/amd64,linux/arm64 -t $(IMAGE) .
