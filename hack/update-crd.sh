#!/usr/bin/env bash

set -eufo pipefail

GO_PROJECT_ROOT="${GO_PROJECT_ROOT:-""}"
CRD_TYPES_PATH="${CRD_TYPES_PATH:-""}"
CRD_OUT_PATH="${CRD_OUT_PATH:-""}"
CRD_FLAG="${CRD_FLAG:-"crd:crdVersions=v1"}"

[ -z "$GO_PROJECT_ROOT" ] && echo "GO_PROJECT_ROOT env var is required" && exit 1
[ -z "$CRD_TYPES_PATH" ] && echo "CRD_TYPES_PATH env var is required" && exit 1
[ -z "$CRD_OUT_PATH" ] && echo "CRD_OUT_PATH env var is required" && exit 1

GO_PROJECT_ROOT=$(realpath ${GO_PROJECT_ROOT})
CRD_TYPES_PATH=$(realpath ${CRD_TYPES_PATH})
CRD_OUT_PATH=$(realpath ${CRD_OUT_PATH})

cd ${GO_PROJECT_ROOT}

# Needs relative paths.
CRD_TYPES_PATH=$(realpath --relative-to="${PWD}" ${CRD_TYPES_PATH})
CRD_OUT_PATH=$(realpath --relative-to="${PWD}" ${CRD_OUT_PATH})

mkdir -p ${CRD_OUT_PATH}
echo "Generating CRD manifests..."

controller-gen \
    "${CRD_FLAG}" \
    paths="./${CRD_TYPES_PATH}/..." \
    output:dir="./${CRD_OUT_PATH}"
