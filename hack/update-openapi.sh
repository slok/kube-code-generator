#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


## CRD paths.
#CRD_PACKAGES=github.com/someone/mypackage/pkg/apis/crdkind/v1alpha1,github.com/someone/mypackage/pkg/apis/crdkind/v1

## Openapi output path.
#OPENAPI_OUTPUT_PACKAGE=github.com/someone/mypackage/openapi

if [[ -z CRD_PACKAGES ]]; then
  echo "CRD_PACKAGES env var is required"
  exit 1
fi

if [[ -z OPENAPI_OUTPUT_PACKAGE ]]; then
  echo "OPENAPI_OUTPUT_PATH env var is required"
  exit 1
fi

openapi-gen \
  -i ${CRD_PACKAGES} \
  -p ${OPENAPI_OUTPUT_PACKAGE}