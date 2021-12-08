#!/usr/bin/env bash

set -eufo pipefail

## Project specific data
#PROJECT_PACKAGE=github.com/slok/kube-code-generator/example
#CLIENT_GENERATOR_OUT=${PROJECT_PACKAGE}/client
#APIS_ROOT=${PROJECT_PACKAGE}/apis
#
## Ugly but needs to be relative if we want to use k8s.io/code-generator
## as it is without touching/sed-ing the code/scripts
#RELATIVE_ROOT_PATH=$(realpath --relative-to="${PWD}" /)
#CODEGEN_PKG=${RELATIVE_ROOT_PATH}${GOPATH}/src/k8s.io/code-generator
#
## Add all groups space separated.
## Example: GROUPS_VERSION="xxxx:v1alpha1 yyyy:v1"
#GROUPS_VERSION="test:v1alpha1"
#
## Generation targets
## Example:
#GENERATION_TARGETS="deepcopy,client"

PROJECT_PACKAGE="${PROJECT_PACKAGE:-""}"
CLIENT_GENERATOR_OUT="${CLIENT_GENERATOR_OUT:-""}"
APIS_ROOT="${APIS_ROOT:-""}"
GROUPS_VERSION="${GROUPS_VERSION:-""}"

[ -z "$PROJECT_PACKAGE" ] && echo "PROJECT_PACKAGE env var is required" && exit 1
[ -z "$CLIENT_GENERATOR_OUT" ] && echo "CLIENT_GENERATOR_OUT env var is required" && exit 1
[ -z "$APIS_ROOT" ] && echo "APIS_ROOT env var is required" && exit 1
[ -z "$GROUPS_VERSION" ] && echo "GROUPS_VERSION env var is required" && exit 1

# Check:
# - https://github.com/kubernetes/community/blob/master/contributors/devel/sig-api-machinery/generating-clientset.md
GENERATION_TARGETS="${GENERATION_TARGETS:-all}"

# From >=1.16 we use gomod, so we need to execute from the project directory.
cd "${GOPATH}/src/${PROJECT_PACKAGE}"

# Ugly but needs to be relative if we want to use k8s.io/code-generator
# as it is without touching/sed-ing the code/scripts
RELATIVE_ROOT_PATH=$(realpath --relative-to="${PWD}" /)
CODEGEN_PKG=${RELATIVE_ROOT_PATH}${GOPATH}/src/k8s.io/code-generator

BOILERPLATE_PATH=/tmp/fake-boilerplate.txt
touch "${BOILERPLATE_PATH}"

# Only generate deepcopy (runtime object needs) and typed client.
# Typed listers & informers not required for the moment. Used with generic
# custom informer/listerwatchers.
${CODEGEN_PKG}/generate-groups.sh \
  ${GENERATION_TARGETS} \
  ${CLIENT_GENERATOR_OUT} \
  ${APIS_ROOT} \
  "${GROUPS_VERSION}" \
  "--go-header-file=${BOILERPLATE_PATH}"
