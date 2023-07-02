IMAGE := ghcr.io/slok/kube-code-generator:v1.27.0

DIRECTORY := $(PWD)
PROJECT_PACKAGE := github.com/slok/kube-code-generator/example
DEPS_CMD := go mod tidy

default: generate

.PHONY: generate
generate: generate-client generate-crd

.PHONY: generate-client
generate-client:
	docker run -it --rm \
	-v $(DIRECTORY):/go/src/$(PROJECT_PACKAGE) \
	-e PROJECT_PACKAGE=$(PROJECT_PACKAGE) \
	-e CLIENT_GENERATOR_OUT=$(PROJECT_PACKAGE)/client \
	-e APIS_ROOT=$(PROJECT_PACKAGE)/apis \
	-e GROUPS_VERSION="comic:v1" \
	-e GENERATION_TARGETS="deepcopy,client" \
	$(IMAGE)

.PHONY: generate-crd
generate-crd:
	docker run -it --rm \
	-v $(DIRECTORY):/src \
	-e GO_PROJECT_ROOT=/src \
	-e CRD_TYPES_PATH=/src/apis \
	-e CRD_OUT_PATH=/src/manifests \
	$(IMAGE) update-crd.sh

.PHONY: deps
deps:
	$(DEPS_CMD)

.PHONY: clean
clean:
	echo "Cleaning generated files..."
	rm -rf ./manifests
	rm -rf ./client
	rm -rf ./apis/comic/v1/zz_generated.deepcopy.go