#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


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


if [[ -z PROJECT_PACKAGE ]]; then
  echo "PROJECT_PACKAGE env var is required"
  exit 1
fi

if [[ -z CLIENT_GENERATOR_OUT ]]; then
  echo "CLIENT_GENERATOR_OUT env var is required"
  exit 1
fi

if [[ -z APIS_ROOT ]]; then
  echo "APIS_ROOT env var is required"
  exit 1
fi

if [[ -z GROUPS_VERSION ]]; then
  echo "GROUPS_VERSION env var is required"
  exit 1
fi

GENERATION_TARGETS="${GENERATION_TARGETS}:-all"

# Ugly but needs to be relative if we want to use k8s.io/code-generator
# as it is without touching/sed-ing the code/scripts
RELATIVE_ROOT_PATH=$(realpath --relative-to="${PWD}" /)
CODEGEN_PKG=${RELATIVE_ROOT_PATH}${GOPATH}/src/k8s.io/code-generator

# Only generate deepcopy (runtime object needs) and typed client.
# Typed listers & informers not required for the moment. Used with generic
# custom informer/listerwatchers.
${CODEGEN_PKG}/generate-groups.sh ${GENERATION_TARGETS} \
  ${CLIENT_GENERATOR_OUT} \
  ${APIS_ROOT} \
  "${GROUPS_VERSION}"