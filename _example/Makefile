IMAGE := ghcr.io/slok/kube-code-generator:latest

DIRECTORY := $(PWD)
DEPS_CMD := go mod tidy

default: generate

.PHONY: generate
generate: generate-client

.PHONY: generate-client
generate-client:
	@docker pull $(IMAGE)
	@docker run -it --rm -v $(DIRECTORY):/app $(IMAGE) \
	--apis-in ./apis \
	--go-gen-out ./client \
	--crd-gen-out ./manifests

.PHONY: deps
deps:
	$(DEPS_CMD)

.PHONY: clean
clean:
	echo "Cleaning generated files..."
	rm -rf ./manifests
	rm -rf ./client
	rm -rf ./apis/comic/v1/zz_generated.deepcopy.go