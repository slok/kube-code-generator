

IMAGE := slok/kube-code-generator:latest

default: build

build:
	docker build -t $(IMAGE) .