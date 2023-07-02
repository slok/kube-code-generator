#!/usr/bin/env sh

set -e

[ -z "$VERSION" ] && echo "VERSION env var is required." && exit 1
[ -z "$IMAGE" ] && echo "IMAGE env var is required." && exit 1
[ -z "$DOCKER_FILE_PATH" ] && echo "DOCKER_FILE_PATH env var is required." && exit 1

IMAGE_TAG="${IMAGE}:${VERSION}"

# Build image.
echo "Building image ${IMAGE_TAG}..."
docker build \
    --build-arg VERSION="${VERSION}" \
    -t "${IMAGE_TAG}" \
    -f "${DOCKER_FILE_PATH}" .
