IMAGE := slok/kube-code-generator:latest

DIRECTORY := $(PWD)
PROJECT_PACKAGE := github.com/slok/kube-code-generator/example

default: generate

generate:
	docker run -it --rm \
	-v $(DIRECTORY):/go/src/$(PROJECT_PACKAGE) \
	-e PROJECT_PACKAGE=$(PROJECT_PACKAGE) \
	-e CLIENT_GENERATOR_OUT=$(PROJECT_PACKAGE)/client \
	-e APIS_ROOT=$(PROJECT_PACKAGE)/apis \
	-e GROUPS_VERSION="test:v1alpha1" \
	-e GENERATION_TARGETS="deepcopy,client" \
	$(IMAGE)
