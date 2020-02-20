#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


## CRD paths.
#CRD_PACKAGES=github.com/someone/mypackage/pkg/apis/crdkind/v1alpha1,github.com/someone/mypackage/pkg/apis/crdkind/v1

## Openapi output path.
#OPENAPI_OUTPUT_PACKAGE=github.com/someone/mypackage/openapi

[ -z "$CRD_PACKAGES" ] && echo "CRD_PACKAGES env var is required" && exit 1;
[ -z "$OPENAPI_OUTPUT_PACKAGE" ] && echo "OPENAPI_OUTPUT_PATH env var is required" && exit 1;

openapi-gen \
  -i ${CRD_PACKAGES} \
  -p ${OPENAPI_OUTPUT_PACKAGE}