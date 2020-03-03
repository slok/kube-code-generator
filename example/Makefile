IMAGE := slok/kube-code-generator:latest

DIRECTORY := $(PWD)
PROJECT_PACKAGE := "github.com/slok/kube-code-generator/example"
K8S_VERSION := "1.17.3"
DEPS_CMD := GO111MODULE=on go mod edit \
	-require=k8s.io/apimachinery@kubernetes-${K8S_VERSION} \
	-require=k8s.io/apiextensions-apiserver@kubernetes-${K8S_VERSION} \
	-require=k8s.io/client-go@kubernetes-${K8S_VERSION} \
	&& go mod tidy

default: generate

.PHONY: generate
generate: generate-client generate-openapi generate-crd

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

.PHONY: generate-openapi
generate-openapi:
	docker run -it --rm \
	-v $(DIRECTORY):/go/src/$(PROJECT_PACKAGE) \
	-e CRD_PACKAGES=$(PROJECT_PACKAGE)/apis/comic/v1 \
	-e OPENAPI_OUTPUT_PACKAGE=$(PROJECT_PACKAGE)/apis/comic/v1 \
	$(IMAGE) update-openapi.sh

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