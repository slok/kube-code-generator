

VERSION ?= $(shell git describe --tags --always)

default: build-image

IMAGE_NAME ?= ghcr.io/slok/kube-code-generator
BUILD_IMAGE_CMD := IMAGE=${IMAGE_NAME} DOCKER_FILE_PATH=./Dockerfile VERSION=${VERSION} ./scripts/build-image.sh
BUILD_PUBLSIH_IMAGE_CMD := IMAGE=${IMAGE_NAME} DOCKER_FILE_PATH=./Dockerfile VERSION=${VERSION} ./scripts/build-publish-image.sh

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

.PHONY: build-image
build-image: ## Builds the docker image.
	@$(BUILD_IMAGE_CMD)

.PHONY: build-publish-image
build-publish-image: ## Builds and publishes docker images.
	@$(BUILD_PUBLSIH_IMAGE_CMD)