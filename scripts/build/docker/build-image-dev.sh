#!/usr/bin/env sh

set -e

[ -z "$VERSION" ] && echo "VERSION env var is required." && exit 1
[ -z "$IMAGE" ] && echo "IMAGE env var is required." && exit 1
[ -z "$DOCKER_FILE_PATH" ] && echo "DOCKER_FILE_PATH env var is required." && exit 1
[ -z "$CW_GO_DEPS_LOGIN" ] && echo "CW_GO_DEPS_LOGIN env var is required (go deps)." && exit 1
[ -z "$CW_GO_DEPS_TOKEN" ] && echo "CW_GO_DEPS_TOKEN env var is required (go deps)." && exit 1

# Build image.
echo "Building dev image ${IMAGE}:${VERSION}..."
docker build \
    --build-arg CW_GO_DEPS_LOGIN="${CW_GO_DEPS_LOGIN}" \
    --build-arg CW_GO_DEPS_TOKEN="${CW_GO_DEPS_TOKEN}" \
    -t "${IMAGE}:${VERSION}" \
    -f "${DOCKER_FILE_PATH}" .
