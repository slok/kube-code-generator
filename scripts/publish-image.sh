#!/usr/bin/env sh

set -e

[ -z "$VERSION" ] && echo "VERSION env var is required." && exit 1
[ -z "$IMAGE" ] && echo "IMAGE env var is required." && exit 1

IMAGE_TAG="${IMAGE}:${VERSION}"

echo "Pushing image ${IMAGE_TAG}..."
docker push ${IMAGE_TAG}
