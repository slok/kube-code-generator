IMAGE := ghcr.io/slok/kube-code-generator:v0.1.0-2-g5fad540-amd64

DIRECTORY := $(PWD)
DEPS_CMD := go mod tidy

default: generate

.PHONY: generate
generate:
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