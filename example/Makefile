IMAGE := slok/kube-code-generator:latest

DIRECTORY := $(PWD)
PROJECT_PACKAGE := github.com/slok/kube-code-generator/example
K8S_VERSION := "1.15.10"
DEPS_CMD := GO111MODULE=on go mod edit \
	-require=k8s.io/apimachinery@kubernetes-${K8S_VERSION} \
	-require=k8s.io/apiextensions-apiserver@kubernetes-${K8S_VERSION} \
	&& go mod tidy

default: generate

.PHONY: generate
generate: generate-client generate-openapi

.PHONY: generate-client
generate-client:
	docker run -it --rm \
	-v $(DIRECTORY):/go/src/$(PROJECT_PACKAGE) \
	-e PROJECT_PACKAGE=$(PROJECT_PACKAGE) \
	-e CLIENT_GENERATOR_OUT=$(PROJECT_PACKAGE)/client \
	-e APIS_ROOT=$(PROJECT_PACKAGE)/apis \
	-e GROUPS_VERSION="test:v1alpha1" \
	-e GENERATION_TARGETS="deepcopy,client" \
	$(IMAGE)

.PHONY: generate-openapi
generate-openapi:
	docker run -it --rm \
	-v $(DIRECTORY):/go/src/$(PROJECT_PACKAGE) \
	-e CRD_PACKAGES=$(PROJECT_PACKAGE)/apis/test/v1alpha1 \
	-e OPENAPI_OUTPUT_PACKAGE=$(PROJECT_PACKAGE)/apis/test/v1alpha1 \
	$(IMAGE) ./update-openapi.sh

.PHONY: deps
deps:
	$(DEPS_CMD)