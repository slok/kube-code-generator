IMAGE := slok/kubernetes-code-generator:latest

DIRECTORY := $(PWD)
PROJECT_PACKAGE := github.com/slok/kube-code-generator/example
UPDATE_CODEGEN_CMD := ./hack/update-codegen.sh

generate: build
	docker run --rm 
	-v $(DIRECTORY):/go/src/$(PROJECT_PACKAGE) \
	$(IMAGE) \
 	/bin/bash -c '$(UPDATE_CODEGEN_CMD)'
