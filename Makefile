

IMAGE := slok/kubernetes-code-generator:latest

build:
	docker build -t $(IMAGE) .