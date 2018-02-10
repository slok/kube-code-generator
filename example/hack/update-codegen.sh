#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Project specific data
PROJECT_PACKAGE=github.com/slok/kube-code-generator/example
CLIENT_GENERATOR_OUT=${PROJECT_PACKAGE}/client
APIS_ROOT=${PROJECT_PACKAGE}/apis

# Ugly but needs to be relative if we want to use k8s.io/code-generator
# as it is without touching/sed-ing the code/scripts
CODEGEN_PKG=./../../../../..${GOPATH}/src/k8s.io/code-generator

# Add all groups space separated.
# Example: GROUPS_VERSION="xxxx:v1alpha1 yyyy:v1"
GROUPS_VERSION="test:v1alpha1"

# Generation targets
# Example: (deepcopy,defaulter,client,lister,informer) or "all".
GENERATION_TARGETS="deepcopy,client"

# Only generate deepcopy (runtime object needs) and typed client.
# Typed listers & informers not required for the moment. Used with generic
# custom informer/listerwatchers.
${CODEGEN_PKG}/generate-groups.sh ${GENERATION_TARGETS} \
  ${CLIENT_GENERATOR_OUT} \
  ${APIS_ROOT} \
  "${GROUPS_VERSION}"