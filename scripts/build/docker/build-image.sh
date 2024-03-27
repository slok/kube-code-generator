#!/usr/bin/env sh

set -e

[ -z "$VERSION" ] && echo "VERSION env var is required." && exit 1
[ -z "$IMAGE" ] && echo "IMAGE env var is required." && exit 1
[ -z "$DOCKER_FILE_PATH" ] && echo "DOCKER_FILE_PATH env var is required." && exit 1
[ -z "$CW_GO_DEPS_LOGIN" ] && echo "CW_GO_DEPS_LOGIN env var is required (go deps)." && exit 1
[ -z "$CW_GO_DEPS_TOKEN" ] && echo "CW_GO_DEPS_TOKEN env var is required (go deps)." && exit 1

# By default use amd64 architecture.
DEF_ARCH=amd64
ARCH=${ARCH:-$DEF_ARCH}

IMAGE_TAG_ARCH="${IMAGE}:${VERSION}-${ARCH}"

# Build image.
echo "Building image ${IMAGE_TAG_ARCH}..."
docker build \
    --build-arg CW_GO_DEPS_LOGIN="${CW_GO_DEPS_LOGIN}" \
    --build-arg CW_GO_DEPS_TOKEN="${CW_GO_DEPS_TOKEN}" \
    --build-arg VERSION="${VERSION}" \
    --build-arg ARCH="${ARCH}" \
    -t "${IMAGE_TAG_ARCH}" \
    -f "${DOCKER_FILE_PATH}" .
